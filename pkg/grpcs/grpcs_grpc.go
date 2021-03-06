// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcs

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	// Sends a greeting
	Create(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

var usersCreateStreamDesc = &grpc.StreamDesc{
	StreamName: "Create",
}

func (c *usersClient) Create(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpcs.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var usersUpdateStreamDesc = &grpc.StreamDesc{
	StreamName: "Update",
}

func (c *usersClient) Update(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpcs.Users/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var usersGetByIdStreamDesc = &grpc.StreamDesc{
	StreamName: "GetById",
}

func (c *usersClient) GetById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpcs.Users/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var usersDeleteStreamDesc = &grpc.StreamDesc{
	StreamName: "Delete",
}

func (c *usersClient) Delete(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpcs.Users/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var usersGetStreamDesc = &grpc.StreamDesc{
	StreamName: "Get",
}

func (c *usersClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/grpcs.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersService is the service API for Users service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUsersService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UsersService struct {
	// Sends a greeting
	Create  func(context.Context, *UserRequest) (*UserResponse, error)
	Update  func(context.Context, *UserRequest) (*UserResponse, error)
	GetById func(context.Context, *UserRequest) (*UserResponse, error)
	Delete  func(context.Context, *UserRequest) (*UserResponse, error)
	Get     func(context.Context, *empty.Empty) (*UsersResponse, error)
}

func (s *UsersService) create(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpcs.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UsersService) update(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpcs.Users/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Update(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UsersService) getById(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpcs.Users/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetById(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UsersService) delete(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpcs.Users/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Delete(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UsersService) get(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpcs.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterUsersService registers a service implementation with a gRPC server.
func RegisterUsersService(s grpc.ServiceRegistrar, srv *UsersService) {
	srvCopy := *srv
	if srvCopy.Create == nil {
		srvCopy.Create = func(context.Context, *UserRequest) (*UserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
		}
	}
	if srvCopy.Update == nil {
		srvCopy.Update = func(context.Context, *UserRequest) (*UserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
		}
	}
	if srvCopy.GetById == nil {
		srvCopy.GetById = func(context.Context, *UserRequest) (*UserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
		}
	}
	if srvCopy.Delete == nil {
		srvCopy.Delete = func(context.Context, *UserRequest) (*UserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
		}
	}
	if srvCopy.Get == nil {
		srvCopy.Get = func(context.Context, *empty.Empty) (*UsersResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpcs.Users",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Create",
				Handler:    srvCopy.create,
			},
			{
				MethodName: "Update",
				Handler:    srvCopy.update,
			},
			{
				MethodName: "GetById",
				Handler:    srvCopy.getById,
			},
			{
				MethodName: "Delete",
				Handler:    srvCopy.delete,
			},
			{
				MethodName: "Get",
				Handler:    srvCopy.get,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "grpcs.proto",
	}

	s.RegisterService(&sd, nil)
}
