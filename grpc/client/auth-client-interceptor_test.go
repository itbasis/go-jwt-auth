package client_test

import (
	"context"
	"errors"
	"flag"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/gofrs/uuid/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testpb"
	"github.com/itbasis/go-jwt-auth/grpc/client"
	"github.com/itbasis/go-jwt-auth/grpc/server"
	"github.com/itbasis/go-jwt-auth/grpc/shared"
	"github.com/itbasis/go-jwt-auth/jwt-token/mocks"
	"github.com/itbasis/go-jwt-auth/model"
	itbasisTestUtils "github.com/itbasis/go-test-utils"
	"github.com/pereslava/grpc_zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthClientInterceptorTestSuite struct {
	*testpb.InterceptorTestSuite

	mockClock    clock.Clock
	mockJwtToken *mocks.JwtToken

	authServerInterceptor *server.AuthServerInterceptor
}

func TestAuthClientInterceptor(t *testing.T) {
	assert.NoError(t, flag.Set("use_tls", "false"))

	testSuite := &AuthClientInterceptorTestSuite{}

	testSuite.mockJwtToken = mocks.NewJwtToken(t)
	testSuite.authServerInterceptor = server.NewAuthServerInterceptorWithCustomParser(testSuite.mockJwtToken)
	testSuite.mockClock = clock.NewMock()

	testSuite.InterceptorTestSuite = &testpb.InterceptorTestSuite{
		TestService: &shared.TestJwtAuthPingService{
			TestServiceServer: &testpb.TestPingService{},
			T:                 t,
		},
		ServerOpts: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				grpc_zerolog.NewUnaryServerInterceptor(itbasisTestUtils.TestLogger),
				auth.UnaryServerInterceptor(testSuite.authServerInterceptor.GetAuthFunc()),
			),
		},
		ClientOpts: []grpc.DialOption{
			grpc.WithUnaryInterceptor(client.NewAuthClientInterceptor().UnaryHeaderAuthorizeForwarder()),
			grpc.WithStreamInterceptor(client.NewAuthClientInterceptor().UnaryStreamHeaderAuthorizeForwarder()),
		},
	}

	suite.Run(t, testSuite)
}

func (s *AuthClientInterceptorTestSuite) BeforeTest(_, _ string) {
	s.mockJwtToken.ExpectedCalls = nil
	s.mockJwtToken.Calls = nil
}

func (s *AuthClientInterceptorTestSuite) AfterTest(_, _ string) {
	assert.True(s.T(), s.mockJwtToken.AssertExpectations(s.T()))
}

func (s *AuthClientInterceptorTestSuite) TestUnary_BadToken() {
	ctx := metadata.AppendToOutgoingContext(s.SimpleCtx(), "Authorization", "bad-token")
	pingResponse, err := s.Client.Ping(ctx, shared.TestGoodPing)
	assert.Error(s.T(), err)

	grpcErr, ok := status.FromError(err)
	assert.True(s.T(), ok)

	assert.Equal(s.T(), codes.Unauthenticated, grpcErr.Code())
	assert.Equal(s.T(), "Bad authorization string", grpcErr.Message())

	assert.Nil(s.T(), pingResponse)
}

func (s *AuthClientInterceptorTestSuite) TestUnary_WithoutToken() {
	pingResponse, err := s.Client.Ping(context.Background(), shared.TestGoodPing)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "something-00000000-0000-0000-0000-000000000000", pingResponse.Value)
}

func (s *AuthClientInterceptorTestSuite) TestUnary_WithToken() {
	s.mockJwtToken.On(
		"Parse", mock.Anything, mock.MatchedBy(
			func(token string) bool {
				return "mock-token" == token
			},
		),
	).Return(&model.SessionUser{UID: uuid.Nil}, nil)
	s.mockJwtToken.On(
		"Parse", mock.Anything, mock.MatchedBy(
			func(token string) bool {
				return "mock-another-token" == token
			},
		),
	).Return(nil, errors.New("bad token"))

	ctx := metadata.AppendToOutgoingContext(s.SimpleCtx(), "Authorization", "bearer mock-token")
	pingResponse, err := s.Client.Ping(ctx, shared.TestGoodPing)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "something-00000000-0000-0000-0000-000000000000", pingResponse.Value)

	ctx = metadata.AppendToOutgoingContext(s.SimpleCtx(), "Authorization", "bearer mock-another-token")
	pingResponse, err = s.Client.Ping(ctx, shared.TestGoodPing)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), pingResponse)
}
