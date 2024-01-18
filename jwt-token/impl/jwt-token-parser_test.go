package impl_test

import (
	"context"

	"github.com/dchest/uniuri"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itbasis/go-clock/v2"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/v2/jwt-token/impl"
	_ "github.com/itbasis/go-test-utils/v2"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
)

var _ = Describe(
	"Parsing token", func() {
		var (
			jwtToken  itbasisJwtToken.JwtToken
			mockClock = clock.NewMock()
			secretKey = "test-key"
		)

		BeforeEach(
			func() {
				var err error

				jwtToken, err = itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, itbasisJwtToken.Config{JwtSecretKey: secretKey})
				Ω(err).Should(Succeed())
			},
		)

		DescribeTable(
			"Invalid token", func(testToken string, expectErr error) {
				Ω(jwtToken.Parse(context.Background(), testToken)).Error().Should(MatchError(ContainSubstring(expectErr.Error())))
			},
			Entry("empty token", "", jwt.ErrTokenMalformed),
			Entry("random string", uniuri.New(), jwt.ErrTokenMalformed),
		)

	},
)
