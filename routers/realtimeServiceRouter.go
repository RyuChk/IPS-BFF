package routers

import (
	"time"

	"github.com/ZecretBone/ips-bff/cmd/realtime-api/handler"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	ginzerolog "github.com/easonlin404/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRealtimeServiceRouter(router *gin.Engine, handlers handler.Handlers) {
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

	oidcConfig := oidcmiddleware.LoadConfig()
	sseV1 := router.Group("/sse/v1")
	{
		sseV1.Use(oidcmiddleware.New(oidcConfig))
		sseV1.GET("/user", handlers.Realtime.OnlineUserBroadcasting)
	}
}
