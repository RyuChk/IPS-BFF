package di

import (
	"git.cie.com/ips/wire-provider/gin/provider"
	"github.com/ZecretBone/ips-bff/cmd/realtime-api/handler"
	"github.com/ZecretBone/ips-bff/cmd/realtime-api/server"
	"github.com/google/wire"
)

var CustomizerSet = wire.NewSet(
	server.ProvideGinRouterCustomizer,
)

var ProviderSet = wire.NewSet(
	handler.ProvideRealtimeHandler,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handlers            *handler.Handlers
	GinServerCustomizer provider.RouterCustomizer
	RealtimeHandler     handler.RealtimeHandler
}
