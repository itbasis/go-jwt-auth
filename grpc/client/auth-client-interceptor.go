package client

import (
	"context"

	grpcMiddlewareMetadata "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
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
	logger := zerolog.Ctx(ctx)

	headerValue := grpcMiddlewareMetadata.ExtractIncoming(ctx).Get(model.HeaderAuthorize)
	logger.Trace().Msgf("headerValue: %s", headerValue)

	if headerValue == "" {
		logger.Trace().Msg("authentication token was not found")

		return ctx
	}

	logger.Trace().Msg("Authentication token found - forward it further in the request...")

	md := metadata.New(map[string]string{model.HeaderAuthorize: headerValue})

	return metadata.NewOutgoingContext(ctx, md)
}
