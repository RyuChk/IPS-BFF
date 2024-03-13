package handler

import (
	"net/http"
	"strconv"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/user/v1"
	mapgrpcclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/mapClient"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	"github.com/gin-gonic/gin"
)

type MapHandler interface {
	GetBuildingList(ctx *gin.Context)
	GetBuildingInfo(ctx *gin.Context)
	GetFloorInfo(ctx *gin.Context)
}

type mapHandler struct {
	mapGRPCSerivce mapgrpcclient.Service
}

func ProvideMapHandler(mapGRPCSerivce mapgrpcclient.Service) MapHandler {
	return &mapHandler{
		mapGRPCSerivce: mapGRPCSerivce,
	}
}

var (
	ToUserV1RoleEnum = map[string]userv1.Role{
		"ADMIN": userv1.Role_ROLE_ADMIN,
		"USER":  userv1.Role_ROLE_USER,
	}
)

func (h *mapHandler) GetBuildingList(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	reqBody := mapv1.GetBuildingListRequest{
		Role: ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)],
	}

	resp, err := h.mapGRPCSerivce.GetBuildingList(ctx, &reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, mapper.ToBuildingModelList(resp))
}

func (h *mapHandler) GetBuildingInfo(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	building := ctx.Param("building")
	reqBody := mapv1.GetBuildingInfoRequest{
		Building: building,
		Role:     ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)],
	}

	resp, err := h.mapGRPCSerivce.GetBuildingDetail(ctx, &reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, mapper.ToBuildingDetailModel(resp))
}

func (h *mapHandler) GetFloorInfo(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	building := ctx.Param("building")
	floor := ctx.Param("floor")
	floor_number, err := strconv.Atoi(floor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "floor number is invalid type")
		return
	}

	reqBody := mapv1.GetFloorInfoRequest{
		Building:    building,
		FloorNumber: int32(floor_number),
		Role:        ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)],
	}

	resp, err := h.mapGRPCSerivce.GetFloorDetail(ctx, &reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, mapper.ToFloorDetailModel(resp))
}
