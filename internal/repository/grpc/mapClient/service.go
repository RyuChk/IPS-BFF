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
	GetBuildingList(ctx context.Context, body *mapv1.GetBuildingListRequest) (*mapv1.GetBuildingListResponse, error)
	GetBuildingDetail(ctx context.Context, body *mapv1.GetBuildingInfoRequest) (*mapv1.GetBuildingInfoResponse, error)
	GetFloorDetail(ctx context.Context, body *mapv1.GetFloorInfoRequest) (*mapv1.GetFloorInfoResponse, error)
}

type mapGRPCClientService struct {
	client mapv1.MapServiceClient
}

func ProvideMapService(config config.GRPCConfig) Service {
	var conn *grpc.ClientConn
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.MapGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (s *mapGRPCClientService) GetBuildingList(ctx context.Context, body *mapv1.GetBuildingListRequest) (*mapv1.GetBuildingListResponse, error) {
	res, err := s.client.GetBuildingList(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *mapGRPCClientService) GetBuildingDetail(ctx context.Context, body *mapv1.GetBuildingInfoRequest) (*mapv1.GetBuildingInfoResponse, error) {
	res, err := s.client.GetBuildingInfo(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *mapGRPCClientService) GetFloorDetail(ctx context.Context, body *mapv1.GetFloorInfoRequest) (*mapv1.GetFloorInfoResponse, error) {
	res, err := s.client.GetFloorInfo(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
