package presenceclient

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

type Service interface {
	FetchOnlineUser(ctx context.Context) (*mapv1.FetchOnlineUserResponse, error)
}

type PresenceGRPCClientService struct {
	client mapv1.PresenceServiceClient
}

func ProvidePresenceServiceClient(config config.GRPCConfig) Service {
	var conn *grpc.ClientConn

	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.RSSIGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		conn = c
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		c, err := grpc.Dial(config.PresenceGRPCHost, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		conn = c
	}

	client := mapv1.NewPresenceServiceClient(conn)
	return &PresenceGRPCClientService{
		client: client,
	}
}

func (s *PresenceGRPCClientService) FetchOnlineUser(ctx context.Context) (*mapv1.FetchOnlineUserResponse, error) {
	resp, err := s.client.FetchOnlineUser(ctx, &mapv1.FetchOnlineUserRequest{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
