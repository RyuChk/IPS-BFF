package rssiclient

import (
	"context"
	"crypto/tls"

	"github.com/ZecretBone/ips-bff/internal/config"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:generate mockgen -source=service.go -destination=mock_rssiclient/mock_service.go -package=mock_rssiclient
type Service interface {
	CollectData(ctx context.Context, body *v1.CollectDataRequest) (*v1.CollectDataResponse, error)
}

type RSSIGRPCClientService struct {
	client v1.StatCollectionServiceClient
}

func ProvideRSSIService(config config.GRPCConfig) Service {
	// Skip verify certificate as the certificate is self-signed
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.Dial(config.RSSIGRPCHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	client := v1.NewStatCollectionServiceClient(conn)
	return &RSSIGRPCClientService{
		client: client,
	}
}

func (s *RSSIGRPCClientService) CollectData(ctx context.Context, body *v1.CollectDataRequest) (*v1.CollectDataResponse, error) {
	res, err := s.client.CollectData(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
