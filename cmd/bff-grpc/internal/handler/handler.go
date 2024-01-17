package handler

import (
	wiregrpc "git.cie.com/ips/wire-provider/grpc"
	bffv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/bff/v1"
)

type Handlers struct {
	RSSI bffv1.BFFServiceServer
}

func RegisterService(server wiregrpc.Server, handler *Handlers) {
	bffv1.RegisterBFFServiceServer(server, handler.RSSI)
}
