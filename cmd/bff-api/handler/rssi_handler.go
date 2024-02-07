package handler

import (
	"fmt"
	"net/http"

	"github.com/ZecretBone/ips-bff/apps/rssi/models/request"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	rssiv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/rssi/v1"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/rssiclient"
	"github.com/ZecretBone/ips-bff/internal/repository/grpc/statclient"
	"github.com/gin-gonic/gin"
)

type RSSIHandler interface {
	GetCoordinate(ctx *gin.Context)
	CollectData(ctx *gin.Context)
	RegisterAp(ctx *gin.Context)
}

type rssiHandler struct {
	rssiStatClient statclient.Service
	rssiClient     rssiclient.Service
}

func ProvideRSSIHandler(rssiStatClient statclient.Service, rssiClient rssiclient.Service) RSSIHandler {
	return &rssiHandler{
		rssiStatClient: rssiStatClient,
		rssiClient:     rssiClient,
	}
}

func (rs *rssiHandler) GetCoordinate(ctx *gin.Context) {
	var body v1.GetCoordinateRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		fmt.Println("theres error in binding json")
	}
	a := ctx.GetHeader("DeviceId")
	b := ctx.GetHeader("Models")

	body.DeviceInfo = &rssiv1.DeviceInfo{
		DeviceId: a,
		Model:    b,
	}

	if _, err := rs.rssiClient.GetCoordinate(ctx, &body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
}

func (rs *rssiHandler) RegisterAp(ctx *gin.Context) {
	var body v1.RegisterApRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		fmt.Println("theres error in binding json")
	}
	fmt.Println("gin")
	fmt.Println(ctx)
	fmt.Println(&body)

	if _, err := rs.rssiClient.RegisterAp(ctx, &body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
}

func (rs *rssiHandler) CollectData(ctx *gin.Context) {
	var body request.StatCollectionRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		ctx.JSON(http.StatusBadRequest, err)
	}

	deviceID := ctx.GetHeader("X-Device-ID")
	model := ctx.GetHeader("X-Device-Model")

	data := mapper.ToDataCollectionDataRequest(body, deviceID, model)

	if _, err := rs.rssiStatClient.CollectData(ctx, data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
