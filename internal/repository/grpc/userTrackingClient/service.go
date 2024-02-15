package usertrackingclient

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

//go:generate mockgen -source=service.go -destination=mock_usertrackingclient/mock_service.go -package=mock_usertrackingclient
type Service interface {
	FetchOnlineUser(ctx context.Context, building string, floor int) (*mapv1.FetchOnlineUserResponse, error)
}

type userTrackingService struct {
	client mapv1.UserTrackingServiceClient
}

func ProvideUserTrackingService(config config.GRPCConfig) Service {
	var conn *grpc.ClientConn
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.UserTrackingGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		conn = c
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		c, err := grpc.Dial(config.UserTrackingGRPCHost, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		conn = c
	}

	client := mapv1.NewUserTrackingServiceClient(conn)
	return &userTrackingService{
		client: client,
	}
}

func (s *userTrackingService) FetchOnlineUser(ctx context.Context, building string, floor int) (*mapv1.FetchOnlineUserResponse, error) {
	resp, err := s.client.FetchOnlineUser(ctx, &mapv1.FetchOnlineUserRequest{
		Building: building,
		Floor:    int32(floor),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
