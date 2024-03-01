package usermanagerclient

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/ZecretBone/ips-bff/internal/config"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=service.go -destination=mock_usermanagerclient/mock_service.go -package=mock_usermanagerclient
type Service interface {
	GetCoordinate(ctx context.Context, body *v1.GetCoordinateRequest) (*v1.GetCoordinateResponse, error)
}

type userManagerGRPCClient struct {
	client v1.UserManagerServiceClient
}

func ProvideUserManagerGRPCClient(config config.GRPCConfig) Service {
	var conn *grpc.ClientConn
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok && currentEnvironment == "beta" {
		c, err := grpc.Dial(config.UserManagerGRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		conn = c
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		c, err := grpc.Dial(config.UserManagerGRPCHost, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		conn = c
	}

	client := v1.NewUserManagerServiceClient(conn)
	return &userManagerGRPCClient{
		client: client,
	}
}

func (s *userManagerGRPCClient) GetCoordinate(ctx context.Context, body *v1.GetCoordinateRequest) (*v1.GetCoordinateResponse, error) {
	res, err := s.client.GetCoordinate(ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
