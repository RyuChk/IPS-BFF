package handler

import (
	"net/http"

	"github.com/ZecretBone/ips-bff/apps/rssi/models/request"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	rssistatclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	"github.com/gin-gonic/gin"
)

type StatCollectionHandler interface {
	CollectData(ctx *gin.Context)
}

type statCollectionHandler struct {
	rssiStatClient rssistatclient.Service
}

func ProvideRSSIHandler(rssiStatClient rssistatclient.Service, rssiClient rssiclient.Service) StatCollectionHandler {
	return &statCollectionHandler{
		rssiStatClient: rssiStatClient,
	}
}

func (rs *statCollectionHandler) CollectData(ctx *gin.Context) {
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
