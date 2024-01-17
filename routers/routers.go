package routers

import (
	"time"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
	ginzerolog "github.com/easonlin404/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine, handlers handler.Handlers) {
	router.Use(ginzerolog.Logger())
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"*", "http://localhost:3000", "http://localhost:8080"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		MaxAge: 12 * time.Hour,
	}))
	rssiV1 := router.Group("/api/v1/rssi")
	{
		rssiV1.POST("/collectdata", handlers.RSSI.CollectData)
		rssiV1.GET("/getcoordinate", handlers.RSSI.GetCoordinate)
		rssiV1.POST("/registerap", handlers.RSSI.RegisterAp)
	}
}
