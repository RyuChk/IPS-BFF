package handler

import (
	"net/http"

	"github.com/ZecretBone/ips-bff/apps/rssi/models/request"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	datacollectionclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/DataCollectionClient"
	"github.com/gin-gonic/gin"
)

type DataCollectionHandler interface {
	CollectData(ctx *gin.Context)
}

type dataCollectionHandler struct {
	statCollectionClient datacollectionclient.Service
}

func ProvideDataCollectionHandler(statCollectionClient datacollectionclient.Service) DataCollectionHandler {
	return &dataCollectionHandler{
		statCollectionClient: statCollectionClient,
	}
}

func (rs *dataCollectionHandler) CollectData(ctx *gin.Context) {
	var body request.StatCollectionRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		ctx.JSON(http.StatusBadRequest, err)
	}

	deviceID := ctx.GetHeader("X-Device-ID")
	model := ctx.GetHeader("X-Device-Model")

	data := mapper.ToDataCollectionDataRequest(body, deviceID, model)

	if _, err := rs.statCollectionClient.CollectData(ctx, data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
