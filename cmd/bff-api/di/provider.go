package di

import (
	"git.cie.com/ips/wire-provider/gin/provider"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/server"
	"github.com/google/wire"
)

var CustomizerSet = wire.NewSet(
	server.ProvideGinRouterCustomizer,
)

var ProviderSet = wire.NewSet(
	handler.ProvideRSSIHandler,
	handler.ProvideMapHandler,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handlers            *handler.Handlers
	GinServerCustomizer provider.RouterCustomizer
	RSSIHandler         handler.RSSIHandler
}
