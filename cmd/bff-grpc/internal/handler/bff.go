package handler

import (
	"context"

	bffv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/bff/v1"
	sharedv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/bff/v1"
	"github.com/ZecretBone/ips-bff/internal/services/rssi"
)

type BFFV1Impl struct {
	rssiService rssi.Service
	bffv1.UnimplementedBFFServiceServer
}

func ProvideBFFServer(rssiService rssi.Service) bffv1.BFFServiceServer {
	return &BFFV1Impl{
		rssiService: rssiService,
	}
}

func (s *BFFV1Impl) HealthCheck(ctx context.Context, req *bffv1.HealthCheckRequest) (*bffv1.HealthCheckResponse, error) {
	return &bffv1.HealthCheckResponse{
		Response: s.rssiService.AddingText(req.TestMessage),
		Status:   sharedv1.HealthStatus_HEALTH_STATUS_OK,
	}, nil
}
