package mapgrpcclient

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/ZecretBone/ips-bff/internal/config"
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=service.go -destination=mock_mapclient/mock_service.go -package=mock_mapclient
type Service interface {
	GetFloorList(ctx context.Context, body *mapv1.GetFloorListRequest) (*mapv1.GetFloorListResponse, error)
	GetFloorMapURL(ctx context.Context, body *mapv1.GetMapURLRequest) (*mapv1.GetMapURLResponse, error)
}

type mapGRPCClientService struct {
	client mapv1.MapServiceClient
}

func ProvideMapService(config config.GRPCConfig) Service {
	var conn *grpc.ClientConn
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.DataCollectionGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		conn = c
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		c, err := grpc.Dial(config.MapGRPCHost, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		conn = c
	}

	client := mapv1.NewMapServiceClient(conn)
	return &mapGRPCClientService{
		client: client,
	}
}

func (s *mapGRPCClientService) GetFloorList(ctx context.Context, body *mapv1.GetFloorListRequest) (*mapv1.GetFloorListResponse, error) {
	res, err := s.client.GetFloorList(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *mapGRPCClientService) GetFloorMapURL(ctx context.Context, body *mapv1.GetMapURLRequest) (*mapv1.GetMapURLResponse, error) {
	res, err := s.client.GetMapURL(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
