package di

import (
	grpcProvide "git.cie.com/ips/wire-provider/grpc/provider"
	"github.com/ZecretBone/ips-bff/cmd/bff-grpc/internal/handler"
	"github.com/ZecretBone/ips-bff/cmd/bff-grpc/server"
	bffv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/bff/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	server.ProvideGRPCServerCustomizer,
	handler.ProvideBFFServer,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handler              *handler.Handlers
	GRPCServerCustomizer grpcProvide.GRPCServerCustomizer
	RSSI                 bffv1.BFFServiceServer
}
