package di

import (
	wiregin "git.cie.com/ips/wire-provider/gin"
	"github.com/ZecretBone/ips-bff/internal/services/realtime"
)

type Container struct {
	server          *wiregin.Server
	RealtimeService realtime.Service
}
