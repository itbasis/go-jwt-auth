package client

import (
	"context"

	grpcMiddlewareMetadata "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthClientInterceptor struct{}

func NewAuthClientInterceptor() *AuthClientInterceptor {
	return &AuthClientInterceptor{}
}

func (receiver *AuthClientInterceptor) UnaryHeaderAuthorizeForwarder() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(receiver.interceptor(ctx), method, req, reply, cc, opts...)
	}
}

func (receiver *AuthClientInterceptor) UnaryStreamHeaderAuthorizeForwarder() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(receiver.interceptor(ctx), desc, cc, method, opts...)
	}
}

func (receiver *AuthClientInterceptor) interceptor(ctx context.Context) context.Context {
	logger := zapctx.Logger(ctx).Sugar()

	headerValue := grpcMiddlewareMetadata.ExtractIncoming(ctx).Get(model.HeaderAuthorize)
	logger.Debugf("headerValue: %s", headerValue)

	if headerValue == "" {
		logger.Debug("authentication token was not found")

		return ctx
	}

	logger.Debug("Authentication token found - forward it further in the request...")

	md := metadata.New(map[string]string{model.HeaderAuthorize: headerValue})

	return metadata.NewOutgoingContext(ctx, md)
}
