package realtime

import (
	"context"
	"io"
	"time"

	"github.com/ZecretBone/ips-bff/apps/rssi/models"
	"github.com/ZecretBone/ips-bff/internal/config"
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
	"github.com/ZecretBone/ips-bff/internal/repository/grpc/presenceclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Service interface {
	StartUserPresenceServer()
	HandleMessageStream(c *gin.Context, username string)
}

type service struct {
	cfg               config.RealtimeServiceConfig
	presenceService   presenceclient.Service
	onlineUserChannel map[string]chan []OnlineUser
}

type OnlineUser struct {
	DisplayName string          `json:"display_name"`
	Timestamp   time.Time       `json:"timestamp"`
	Coordinate  models.Position `json:"coordinate"`
}

func ProvideRealtimeService(cfg config.RealtimeServiceConfig, presenceService presenceclient.Service) Service {
	return &service{
		cfg:               cfg,
		presenceService:   presenceService,
		onlineUserChannel: make(map[string]chan []OnlineUser),
	}
}

func mapFetchResponseToOnlineUserStruct(resp *mapv1.FetchOnlineUserResponse) []OnlineUser {
	result := make([]OnlineUser, len(resp.OnlineUsers))
	for i, v := range resp.OnlineUsers {
		result[i] = OnlineUser{
			DisplayName: v.DisplayName,
			Coordinate: models.Position{
				X: float64(v.Coordinate.X),
				Y: float64(v.Coordinate.Y),
				Z: float64(v.Coordinate.Z),
			},
			Timestamp: v.Timestamp.AsTime(),
		}
	}

	return result
}

func (s *service) StartUserPresenceServer() {
	log.Info().Str("polling_rate", s.cfg.OnlineUserPresenceDelay.String()).Msg("Starting presence polling connection")
	for {
		resp, err := s.presenceService.FetchOnlineUser(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("error while fetching online user from presence service")
			continue
		}
		result := mapFetchResponseToOnlineUserStruct(resp)

		for _, v := range s.onlineUserChannel {
			v <- result
		}
		time.Sleep(s.cfg.OnlineUserPresenceDelay)
	}
}

func (s *service) HandleMessageStream(c *gin.Context, username string) {
	s.onlineUserChannel[username] = make(chan []OnlineUser)
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-s.onlineUserChannel[username]; ok {
			c.SSEvent("online_user", msg)
			return true
		}
		return false
	})
}
