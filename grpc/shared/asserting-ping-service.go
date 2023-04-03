package shared

import (
	"context"
	"fmt"
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testpb"
)

var (
	TestGoodPing = &testpb.PingRequest{Value: "something", SleepTimeMs: 9999} //nolint:gomnd
)

type TestJwtAuthPingService struct {
	testpb.TestServiceServer
	T *testing.T
}

func (s *TestJwtAuthPingService) Ping(ctx context.Context, ping *testpb.PingRequest) (*testpb.PingResponse, error) {
	ping.Value = fmt.Sprintf("%s-%s", ping.Value, getSessionUserID(ctx))

	return s.TestServiceServer.Ping(ctx, ping) //nolint:wrapcheck
}

func (s *TestJwtAuthPingService) PingError(ctx context.Context, ping *testpb.PingErrorRequest) (*testpb.PingErrorResponse, error) {
	ping.Value = fmt.Sprintf("%s-%s", ping.Value, getSessionUserID(ctx))

	return s.TestServiceServer.PingError(ctx, ping) //nolint:wrapcheck
}

func getSessionUserID(ctx context.Context) uuid.UUID {
	sessionUser, err := GetSessionUser(ctx)
	if err != nil {
		return uuid.Nil
	}

	return sessionUser.UID
}
