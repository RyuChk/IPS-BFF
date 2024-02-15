package handler

import (
	"net/http"
	"strconv"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	usertrackingclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/userTrackingClient"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	"github.com/gin-gonic/gin"
)

type UserTrackingHandler interface {
	GetOnlineUser(ctx *gin.Context)
}

type userTrackingHandler struct {
	userTrackingClient usertrackingclient.Service
}

func ProvideUserTrackingHandler(userTrackingClient usertrackingclient.Service) UserTrackingHandler {
	return &userTrackingHandler{
		userTrackingClient: userTrackingClient,
	}
}

func (h *userTrackingHandler) GetOnlineUser(ctx *gin.Context) {
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

	building := ctx.Param("building")
	floor, err := strconv.Atoi(ctx.Param("floor"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "floor input is invalid type")
		return
	}

	resp, err := h.userTrackingClient.FetchOnlineUser(ctx, building, floor)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, mapper.MapFetchResponseToOnlineUserStruct(resp))
}
