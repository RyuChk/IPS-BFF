package handler

import (
	"net/http"
	"time"

	"github.com/ZecretBone/ips-bff/internal/services/realtime"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	"github.com/gin-gonic/gin"
)

type RealtimeHandler interface {
	OnlineUserBroadcasting(ctx *gin.Context)
}

type realtimeHandler struct {
	realtimeService realtime.Service
}

func ProvideRealtimeHandler(realtimeService realtime.Service) RealtimeHandler {
	return &realtimeHandler{
		realtimeService: realtimeService,
	}
}

func (h *realtimeHandler) OnlineUserBroadcasting(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		ctx.Abort()
		return
	}

	if !userInfo.IsAdmin() {
		ctx.JSON(http.StatusForbidden, "Your role is not eligible for this function")
		ctx.Abort()
		return
	}
	h.realtimeService.HandleMessageStream(ctx, time.Now().String())
}
