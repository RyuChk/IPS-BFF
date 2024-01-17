package handler

import (
	"fmt"
	"net/http"

	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	rssiv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/rssi/v1"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/rssi"
	rssistatclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	"github.com/gin-gonic/gin"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

type RSSIHandler interface {
	GetCoordinate(ctx *gin.Context)
	CollectData(ctx *gin.Context)
	RegisterAp(ctx *gin.Context)
}

type rssiHandler struct {
	rssiStatClient rssistatclient.Service
	rssiClient     rssiclient.Service
}

func ProvideRSSIHandler(rssiStatClient rssistatclient.Service, rssiClient rssiclient.Service) RSSIHandler {
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
		Models:   b,
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
	var body v1.CollectDataRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		fmt.Println("theres error in binding json")
	}
	body.Stage = rssiv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE
	a := ctx.GetHeader("DeviceId")
	b := ctx.GetHeader("Models")

	body.DeviceInfo = &rssiv1.DeviceInfo{
		DeviceId: a,
		Models:   b,
	}
	// body := v1.CollectDataRequest{
	// 	Signals: []*rssiv1.RSSI{
	// 		{
	// 			Ssid:        "newtONE",
	// 			MacAddress:  "ab:cd:ef:gh:xy:xz",
	// 			Strength:    60.3,
	// 			PollingRate: 300,
	// 			CreatedAt:   timestamppb.Now(),
	// 		}, {
	// 			Ssid:        "newtTWO",
	// 			MacAddress:  "ab:cd:ef:gh:xy:x1",
	// 			Strength:    50.3,
	// 			PollingRate: 300,
	// 			CreatedAt:   timestamppb.Now(),
	// 		},
	// 	},
	// 	DeviceInfo: &rssiv1.DeviceInfo{
	// 		DeviceId: "FXM1234",
	// 		Models:   "androidOreo123",
	// 	},
	// 	Stage: rssiv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE,
	// 	Position: &rssiv1.Position{
	// 		X: 3,
	// 		Y: 2,
	// 		Z: 1,
	// 	},
	// 	Duration:  10,
	// 	StartedAt: timestamppb.Now(),
	// 	EndedAt:   timestamppb.Now(),
	// 	CreatedAt: timestamppb.Now(),
	// }
	if _, err := rs.rssiStatClient.CollectData(ctx, &body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
}
