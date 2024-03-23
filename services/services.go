package services

import (
	"fmt"

	"EXAM3/api-gateway/config"
	pbp "EXAM3/api-gateway/genproto/product_service"
	pb "EXAM3/api-gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pb.UserServiceClient
	ProductService() pbp.ProductServiceClient
}

type serviceManager struct {
	userService    pb.UserServiceClient
	productService pbp.ProductServiceClient
}

func (s *serviceManager) UserService() pb.UserServiceClient {
	return s.userService
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connProduct, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProductServiceHost, conf.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pb.NewUserServiceClient(connUser),
		productService: pbp.NewProductServiceClient(connProduct),
	}

	return serviceManager, nil
}
