package rssiclient

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/ZecretBone/ips-bff/internal/config"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
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
	var conn *grpc.ClientConn
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.RSSIGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		conn = c
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		c, err := grpc.Dial(config.RSSIGRPCHost, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		conn = c
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
