//go:build wireinject
// +build wireinject

package di

import (
	"git.cie-ips.com/ips/wire-provider/gin/provider"
	internalDi "github.com/ZecretBone/ips-bff/internal/di"
	"github.com/google/wire"
)

var BaseBindingSet = wire.NewSet(
	ProviderSet,
	CustomizerSet,

	internalDi.DatabaseSet,
	internalDi.ConfigSet,
	internalDi.ProviderSet,
)

var MainBindingSet = wire.NewSet(
	BaseBindingSet,
	provider.WireSet,
	wire.Struct(new(Container), "*"),
)

func InitializeContainer() (*Container, func(), error) {
	wire.Build(MainBindingSet)
	return &Container{}, func() {}, nil
}
