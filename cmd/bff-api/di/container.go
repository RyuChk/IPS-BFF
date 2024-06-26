package di

import (
	wiregin "git.cie-ips.com/ips/wire-provider/gin"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
)

type Container struct {
	server          *wiregin.Server
	UserTrackingSSE handler.UserTrackingSSEHandler
}
