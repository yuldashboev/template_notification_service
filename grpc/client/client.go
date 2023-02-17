package client

import (
	"context"
	"fmt"

	pbgorush "github.com/appleboy/gorush/rpc/proto"
	"google.golang.org/grpc"

	pbc "genproto/courier_service"
	pbu "genproto/user_service"

	"gitlab.udevs.io/template/template_notification_service/config"
)

type GrpcClientI interface {
	CourierService() pbc.CourierServiceClient
	BranchUserService() pbu.BranchUserServiceClient
	GorushService() pbgorush.GorushClient
	CustomerService() pbu.CustomerServiceClient
}

type grpcClient struct {
	cfg               config.Config
	courierService    pbc.CourierServiceClient
	branchUserService pbu.BranchUserServiceClient
	gorushService     pbgorush.GorushClient
	customerService   pbu.CustomerServiceClient
}

func New(ctx context.Context, cfg config.Config) (*grpcClient, error) {
	connCourier, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%s", cfg.CourierServiceHost, cfg.CourierServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port: %s",
			cfg.CourierServiceHost, cfg.CourierServicePort)
	}

	connGorush, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%s", cfg.GorushHost, cfg.GorushPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("gorush service dial host: %s port: %s",
			cfg.GorushHost, cfg.GorushPort)
	}

	connUser, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port: %s",
			cfg.UserServiceHost, cfg.UserServicePort)
	}

	return &grpcClient{
		cfg:               cfg,
		courierService:    pbc.NewCourierServiceClient(connCourier),
		branchUserService: pbu.NewBranchUserServiceClient(connUser),
		gorushService:     pbgorush.NewGorushClient(connGorush),
		customerService:   pbu.NewCustomerServiceClient(connUser),
	}, nil
}

func (g *grpcClient) CourierService() pbc.CourierServiceClient {
	return g.courierService
}

func (g *grpcClient) GorushService() pbgorush.GorushClient {
	return g.gorushService
}

func (g *grpcClient) BranchUserService() pbu.BranchUserServiceClient {
	return g.branchUserService
}

func (g *grpcClient) CustomerService() pbu.CustomerServiceClient {
	return g.customerService
}
