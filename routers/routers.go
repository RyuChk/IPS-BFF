package routers

import (
	"time"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
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
		rssiV1.POST("/collectdata", handlers.StatCollection.CollectData)
	}

	oidcConfig := oidcmiddleware.LoadConfig()

	userManagerV1 := router.Group("/api/v1/user")
	{
		userManagerV1.Use(oidcmiddleware.New(oidcConfig))

		userManagerV1.GET("/ws", handlers.UserManager.GetCoordinate)

		admin := userManagerV1.Group("/admin")
		{
			admin.GET("/online/:building/:floor", handlers.UserTracking.GetOnlineUser)
			admin.GET("/online/sse/:building/:floor", handlers.UserTrackingSSE.HandleNewClient)
		}
	}

	mapv1 := router.Group("/api/v1/map")
	{
		mapv1.Use(oidcmiddleware.New(oidcConfig))

		mapv1.GET("/building", handlers.Map.GetBuildingList)
		mapv1.GET("/info/:building", handlers.Map.GetBuildingInfo)
		mapv1.GET("/info/:building/:floor", handlers.Map.GetFloorInfo)
	}
}
