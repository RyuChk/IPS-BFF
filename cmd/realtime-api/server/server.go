package server

import (
	wiregin "git.cie.com/ips/wire-provider/gin"
	"git.cie.com/ips/wire-provider/gin/provider"

	"github.com/ZecretBone/ips-bff/cmd/realtime-api/handler"
	"github.com/ZecretBone/ips-bff/routers"
)

type ginRouterCustomizer struct {
	handler *handler.Handlers
}

func (gc *ginRouterCustomizer) Register(server *wiregin.Server) error {
	routers.RegisterRealtimeServiceRouter(server.Engine, *gc.handler)
	return nil
}

func (gc *ginRouterCustomizer) Configure(builder wiregin.Builder) error {
	return nil
}

func ProvideGinRouterCustomizer(handler handler.Handlers) provider.RouterCustomizer {
	return &ginRouterCustomizer{
		handler: &handler,
	}
}
