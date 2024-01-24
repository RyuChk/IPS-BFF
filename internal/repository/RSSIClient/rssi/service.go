package rssiclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/ZecretBone/ips-bff/internal/config"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=service.go -destination=mock_rssiclient/mock_service.go -package=mock_rssiclient
type Service interface {
	GetCoordinate(ctx context.Context, body *v1.GetCoordinateRequest) (*v1.GetCoordinateResponse, error)
	RegisterAp(ctx context.Context, body *v1.RegisterApRequest) (*v1.RegisterApResponse, error)
}

type RSSIGRPCClientService struct {
	client v1.CoordinateCollectionServiceClient
}

func ProvideRSSIService(config config.GRPCConfig) Service {
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

	client := v1.NewCoordinateCollectionServiceClient(conn)
	return &RSSIGRPCClientService{
		client: client,
	}
}

func (s *RSSIGRPCClientService) GetCoordinate(ctx context.Context, body *v1.GetCoordinateRequest) (*v1.GetCoordinateResponse, error) {
	res, err := s.client.GetCoordinate(ctx, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RSSIGRPCClientService) RegisterAp(ctx context.Context, body *v1.RegisterApRequest) (*v1.RegisterApResponse, error) {
	fmt.Println("mybody: ")
	fmt.Println(body)
	res, err := s.client.RegisterAp(ctx, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
