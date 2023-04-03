package impl_test

import (
	"context"

	"github.com/benbjohnson/clock"
	"github.com/dchest/uniuri"
	"github.com/golang-jwt/jwt/v5"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/jwt-token/impl"
	itbasisTestUtils "github.com/itbasis/go-test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Parsing with secret key", func() {
		var ctx context.Context
		var jwtToken itbasisJwtToken.JwtToken
		mockClock := clock.NewMock()
		secretKey := "test-key"

		BeforeEach(
			func() {
				ctx = itbasisTestUtils.TestLoggerWithContext(context.Background())
				var err error

				jwtToken, err = itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, itbasisJwtToken.Config{JwtSecretKey: secretKey})
				Ω(err).Should(Succeed())
			},
		)

		DescribeTable(
			"Invalid token", func(testToken string, expectErr error) {
				sessionUser, err := jwtToken.Parse(ctx, testToken)
				Ω(err).Should(MatchError(expectErr))
				Ω(sessionUser).To(BeNil())
			},
			Entry("empty token", "", jwt.ErrTokenMalformed),
			Entry("random string", uniuri.New(), jwt.ErrTokenMalformed),
		)
	},
)
