package grpcserver

import (
	"context"
	"github.com/dalmarcogd/gbpl-go/internal/models"
	"github.com/dalmarcogd/gbpl-go/pkg/grpcs"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ServiceImpl) CreateUser(ctx context.Context, c *grpcs.UserRequest) (*grpcs.UserResponse, error) {
	user := models.User{
		Name:  &c.Name,
		Email: &c.Email,
	}
	err := s.ServiceManager().Handlers().CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &grpcs.UserResponse{
		Id:    user.Id,
		Name:  *user.Name,
		Email: *user.Email,
	}, nil
}

func (s *ServiceImpl) UpdateUser(ctx context.Context, c *grpcs.UserRequest) (*grpcs.UserResponse, error) {
	user := models.User{
		Id:    c.Id,
		Name:  &c.Name,
		Email: &c.Email,
	}
	err := s.ServiceManager().Handlers().UpdateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &grpcs.UserResponse{
		Id:    user.Id,
		Name:  *user.Name,
		Email: *user.Email,
	}, nil
}

func (s *ServiceImpl) GetUserById(ctx context.Context, c *grpcs.UserRequest) (*grpcs.UserResponse, error) {
	user := models.User{
		Id: c.Id,
	}
	err := s.ServiceManager().Handlers().GetUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &grpcs.UserResponse{
		Id:    user.Id,
		Name:  *user.Name,
		Email: *user.Email,
	}, nil
}

func (s *ServiceImpl) GetUsers(ctx context.Context, _ *empty.Empty) (*grpcs.UsersResponse, error) {
	var users []models.User
	err := s.ServiceManager().Handlers().GetUsers(ctx, &users)
	if err != nil {
		return nil, err
	}

	uResponses := make([]*grpcs.UserResponse, 0)
	for _, user := range users {
		uResponses = append(uResponses, &grpcs.UserResponse{
			Id:    user.Id,
			Name:  *user.Name,
			Email: *user.Email,
		})
	}
	return &grpcs.UsersResponse{Users: uResponses}, nil
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, request *grpcs.UserRequest) (*grpcs.UserResponse, error) {
	user := models.User{
		Id: request.Id,
	}

	err := s.ServiceManager().Handlers().DeleteUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &grpcs.UserResponse{
		Id:    user.Id,
		Name:  *user.Name,
		Email: *user.Email,
	}, nil
}
