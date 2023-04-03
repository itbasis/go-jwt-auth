package server_test

import (
	"flag"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/gofrs/uuid/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testpb"
	jwtAuthGrpcServer "github.com/itbasis/go-jwt-auth/grpc/server"
	jwtAuthGrpcShared "github.com/itbasis/go-jwt-auth/grpc/shared"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	itbasisJwtAuthMocks "github.com/itbasis/go-jwt-auth/jwt-token/mocks"
	itbasisJwtAuthModel "github.com/itbasis/go-jwt-auth/model"
	testUtils "github.com/itbasis/go-test-utils"
	"github.com/pereslava/grpc_zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthServerInterceptorTestSuite struct {
	*testpb.InterceptorTestSuite

	mockClock    clock.Clock
	mockJwtToken *itbasisJwtAuthMocks.JwtToken

	authServerInterceptor *jwtAuthGrpcServer.AuthServerInterceptor
}

func TestAuthServerInterceptor(t *testing.T) {
	assert.NoError(t, flag.Set("use_tls", "false"))

	testSuite := &AuthServerInterceptorTestSuite{}

	testSuite.mockJwtToken = itbasisJwtAuthMocks.NewJwtToken(t)
	testSuite.authServerInterceptor = jwtAuthGrpcServer.NewAuthServerInterceptorWithCustomParser(testSuite.mockJwtToken)
	testSuite.mockClock = clock.NewMock()

	testSuite.InterceptorTestSuite = &testpb.InterceptorTestSuite{
		TestService: &jwtAuthGrpcShared.TestJwtAuthPingService{
			TestServiceServer: &testpb.TestPingService{},
			T:                 t,
		},
		ServerOpts: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				grpc_zerolog.NewUnaryServerInterceptor(testUtils.TestLogger),
				auth.UnaryServerInterceptor(testSuite.authServerInterceptor.GetAuthFunc()),
			),
			grpc.ChainStreamInterceptor(
				auth.StreamServerInterceptor(testSuite.authServerInterceptor.GetAuthFunc()),
			),
		},
	}

	suite.Run(t, testSuite)
}

func (s *AuthServerInterceptorTestSuite) BeforeTest(_, _ string) {
	s.mockJwtToken.ExpectedCalls = nil
	s.mockJwtToken.Calls = nil
}
func (s *AuthServerInterceptorTestSuite) AfterTest(_, _ string) {
	assert.True(s.T(), s.mockJwtToken.AssertExpectations(s.T()))
}

func (s *AuthServerInterceptorTestSuite) TestUnary_NoAuth() {
	_, err := s.Client.Ping(s.SimpleCtx(), jwtAuthGrpcShared.TestGoodPing)
	assert.NoError(s.T(), err)
	s.mockJwtToken.AssertNumberOfCalls(s.T(), "Parse", 0)
}

func (s *AuthServerInterceptorTestSuite) TestUnary_InvalidToken() {
	tests := []struct {
		headerValue      string
		expectErrorMsg   string
		expectCallParser int
	}{
		{headerValue: "test", expectErrorMsg: "Bad authorization string"},
		{headerValue: "Bearer test", expectErrorMsg: "authentication required", expectCallParser: 1},
	}
	for _, test := range tests {
		s.T().Run(
			test.headerValue, func(t *testing.T) {
				s.mockJwtToken.
					On("Parse", mock.Anything, mock.Anything).
					Return(nil, itbasisJwtToken.ErrTokenInvalid)

				ctx := metadata.AppendToOutgoingContext(s.SimpleCtx(), "authorization", test.headerValue)
				_, err := s.Client.Ping(ctx, jwtAuthGrpcShared.TestGoodPing)

				assert.Error(t, err)

				grpcErr, ok := status.FromError(err)
				assert.True(t, ok)

				assert.Equal(t, codes.Unauthenticated, grpcErr.Code())
				assert.Equal(t, test.expectErrorMsg, grpcErr.Message())

				s.mockJwtToken.AssertNumberOfCalls(s.T(), "Parse", test.expectCallParser)
			},
		)
	}
}

func (s *AuthServerInterceptorTestSuite) TestUnary_Success() {
	s.mockJwtToken.
		On("Parse", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				testUtils.TestLogger.Info().Msgf("args: %++v", args)
			},
		).
		Return(&itbasisJwtAuthModel.SessionUser{UID: uuid.Nil}, nil)

	ctx := metadata.AppendToOutgoingContext(s.SimpleCtx(), "authorization", "bearer mock-token")
	_, err := s.Client.Ping(ctx, jwtAuthGrpcShared.TestGoodPing)

	assert.NoError(s.T(), err)
	s.mockJwtToken.AssertNumberOfCalls(s.T(), "Parse", 1)
}
