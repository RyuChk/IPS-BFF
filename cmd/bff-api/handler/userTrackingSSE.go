package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	"github.com/ZecretBone/ips-bff/internal/models"
	usertrackingclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/userTrackingClient"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserTrackingSSEHandler interface {
	HandleNewClient(c *gin.Context)
}

type userTrackingSSEHandler struct {
	BroadcastChannel map[string]map[string]chan []models.OnlineUser
	FetchingKey      map[string]int

	userTrackingClient usertrackingclient.Service
}

func ProvideUserTrackingSSEHandler(userTrackingClient usertrackingclient.Service) UserTrackingSSEHandler {
	return &userTrackingSSEHandler{
		BroadcastChannel:   make(map[string]map[string]chan []models.OnlineUser),
		FetchingKey:        make(map[string]int),
		userTrackingClient: userTrackingClient,
	}
}

func (h *userTrackingSSEHandler) HandleNewClient(c *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(c)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		c.Abort()
		return
	}

	if !userInfo.IsAdmin() {
		c.JSON(http.StatusForbidden, "Your role is not eligible for this function")
		c.Abort()
		return
	}

	key := fmt.Sprintf("%s:%s", c.Param("building"), c.Param("floor"))
	uid := time.Now().String()
	log.Info().Str("key", key).Str("uid", uid).Msg("user connected")
	if _, exist := h.BroadcastChannel[key]; !exist {
		log.Debug().Str("key", key).Str("uid", uid).Msg("start new broadcasting channel")

		h.FetchingKey[key]++
		h.BroadcastChannel[key] = make(map[string]chan []models.OnlineUser)
		h.BroadcastChannel[key][uid] = make(chan []models.OnlineUser)
		go h.HandleUpdateChannel(key)
	} else {
		log.Debug().Str("key", key).Str("uid", uid).Msg("use exist channel")

		h.FetchingKey[key]++
		h.BroadcastChannel[key][uid] = make(chan []models.OnlineUser)
	}

	defer h.HandleCloseClient(key, uid)

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-h.BroadcastChannel[key][uid]; ok {
			c.SSEvent("online_user", msg)
			return true
		}
		return false
	})
}

func (h *userTrackingSSEHandler) HandleUpdateChannel(key string) {
	for h.FetchingKey[key] > 0 {
		params := strings.Split(key, ":")
		floor, err := strconv.Atoi(params[1])
		if err != nil {
			log.Error().Err(err).Msg("error conv params")
			continue
		}

		resp, err := h.userTrackingClient.FetchOnlineUser(context.Background(), params[0], floor)
		if err != nil {
			log.Error().Err(err).Str("key", key).Msg("Error fetching online values from current key")
			continue
		}

		for _, ch := range h.BroadcastChannel[key] {
			ch <- mapper.MapFetchResponseToOnlineUserStruct(resp)
		}
		time.Sleep(5 * time.Second)
	}

	log.Debug().Str("key", key).Msg("delete channel; no user currently in this channel")
	delete(h.BroadcastChannel, key)
}

func (h *userTrackingSSEHandler) HandleCloseClient(key, uid string) {
	log.Info().Str("key", key).Str("uid", uid).Msg("user disconnected")

	h.FetchingKey[key]--
	delete(h.BroadcastChannel[key], uid)
}
