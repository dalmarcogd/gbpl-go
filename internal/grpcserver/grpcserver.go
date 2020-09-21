package grpcserver

import (
	"context"
	"github.com/dalmarcogd/gbpl-go/internal/services"
	"github.com/dalmarcogd/gbpl-go/pkg/grpcs"
	"google.golang.org/grpc"
	"net"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		grpcServer     *grpc.Server
		address        string
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithAddress(address string) *ServiceImpl {
	s.address = address
	return s
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.grpcServer = grpc.NewServer()
	s.RegisterServices()
	return nil
}

func (s *ServiceImpl) Close() error {
	s.grpcServer.GracefulStop()
	return nil
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.GrpcServer {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) RegisterServices() *ServiceImpl {
	grpcs.RegisterUsersService(s.grpcServer, &grpcs.UsersService{
		Create:  s.CreateUser,
		Update:  s.UpdateUser,
		GetById: s.GetUserById,
		Get:     s.GetUsers,
		Delete:  s.DeleteUser,
	})
	return s
}

func (s *ServiceImpl) Run() error {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(listen)
}
