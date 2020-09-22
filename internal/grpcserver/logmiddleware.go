package grpcserver

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/gbpl-go/internal/services"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path"
	"time"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds kit.Logger to the context.
func LogUnaryServerInterceptor(s services.ServiceManager) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		f := grpc_ctxtags.Extract(ctx).Values()
		service := path.Dir(info.FullMethod)[1:]
		method := path.Base(info.FullMethod)
		f["grpc.method"] = method
		f["grpc.service"] = service
		if d, ok := ctx.Deadline(); ok {
			f["grpc.request.deadline"] = d.Format(time.RFC3339)
		}
		s.Logger().Info(ctx, fmt.Sprintf("Request %v", info.FullMethod), f)
		resp, err := handler(ctx, req)
		endTime := time.Now()

		f["grpc.duration"] = endTime.Sub(startTime).String()
		var code codes.Code
		fromError, ok := status.FromError(err)
		if ok {
			code = fromError.Code()
		}
		f["grpc.code"] = code

		s.Logger().Info(ctx, fmt.Sprintf("Response %v code=%v", info.FullMethod, code.String()), f)
		return resp, err
	}
}
