package jwttoken_test

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itbasis/go-clock"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/v2/jwt-token/impl"
	itbasisJwtAuthModel "github.com/itbasis/go-jwt-auth/v2/model"
	testUtils "github.com/itbasis/go-test-utils/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"JwtToken", func() {
		var (
			jwtTokenConfig = itbasisJwtToken.Config{
				JwtSecretKey:                     "test-key",
				JwtSigningMethod:                 "HS512",
				JwtAccessTokenDurationInSeconds:  itbasisJwtToken.TokenDuration(5 * time.Second),
				JwtRefreshTokenDurationInSeconds: itbasisJwtToken.TokenDuration(5 * time.Second),
			}
			mockClock = clock.NewMock()
			logger    = testUtils.TestLogger.Sugar()
		)

		BeforeEach(
			func() {
				mockClock.Set(time.Now())
			},
		)

		Context(
			"Success creating access token", func() {
				var sessionUserTestData = []itbasisJwtAuthModel.SessionUser{
					{Username: "test-user"},
					{UID: uuid.FromStringOrNil("39910003-f693-44ae-979f-f83714a6d459"), Username: "test-user", Email: "test@example.org"},
				}

				for i, testSessionUser := range sessionUserTestData {
					It(
						fmt.Sprintf("Test #%d", i), func() {
							ctx := context.Background()

							jwtToken, err := itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, jwtTokenConfig)
							Ω(err).Should(Succeed())

							accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
							Ω(err).Should(Succeed())
							Ω(accessToken).NotTo(BeEmpty())

							Ω(jwtToken.Parse(ctx, accessToken)).To(HaveValue(Equal(testSessionUser)))
						},
					)
				}
			},
		)

		DescribeTable(
			"Empty fields", func(testSessionUser itbasisJwtAuthModel.SessionUser, expectErr error) {
				ctx := context.Background()

				jwtToken, err := itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, jwtTokenConfig)
				Ω(err).Should(Succeed())

				accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
				Ω(accessToken).NotTo(BeEmpty())
				Ω(err).Should(Succeed())
				logger.Infof("accessToken: %s", accessToken)

			},
			Entry(
				"empty email", itbasisJwtAuthModel.SessionUser{
					UID:      uuid.FromStringOrNil("39910003-f693-44ae-979f-f83714a6d459"),
					Username: "test-user",
				}, itbasisJwtToken.ErrTokenInvalidEmail,
			),
		)

		Context(
			"Empty fields", func() {
				It(
					"empty email", func() {
						ctx := context.Background()

						jwtToken, err := itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, jwtTokenConfig)
						Ω(err).Should(Succeed())

						testSessionUser := itbasisJwtAuthModel.SessionUser{
							UID:      uuid.FromStringOrNil("39910003-f693-44ae-979f-f83714a6d459"),
							Username: "test-user",
						}

						accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
						Ω(accessToken).NotTo(BeEmpty())
						Ω(err).Should(Succeed())
						logger.Infof("accessToken: %s", accessToken)

						Ω(jwtToken.Parse(ctx, accessToken)).Error().Should(MatchError(itbasisJwtToken.ErrTokenInvalidEmail))
					},
				)

				It(
					"empty UID", func() {
						ctx := context.Background()

						jwtToken, err := itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, jwtTokenConfig)
						Ω(err).Should(Succeed())

						testSessionUser := itbasisJwtAuthModel.SessionUser{
							Username: "test-user",
						}

						accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
						Ω(accessToken).NotTo(BeEmpty())
						Ω(err).Should(Succeed())
						logger.Infof("accessToken: %s", accessToken)

						Ω(jwtToken.Parse(ctx, accessToken)).
							Error().Should(
							SatisfyAll(
								MatchError(ContainSubstring(itbasisJwtToken.ErrTokenInvalidUID.Error())),
							),
						)
					},
				)

				It(
					"empty issuer", func() {
						ctx := context.Background()

						jwtToken, err := itbasisJwtTokenImpl.NewJwtTokenCustomConfig(mockClock, jwtTokenConfig)
						Ω(err).Should(Succeed())

						testSessionUser := itbasisJwtAuthModel.SessionUser{
							UID: uuid.FromStringOrNil("39910003-f693-44ae-979f-f83714a6d459"),
						}

						accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
						Ω(accessToken).NotTo(BeEmpty())
						Ω(err).Should(Succeed())
						logger.Infof("accessToken: %s", accessToken)

						Ω(jwtToken.Parse(ctx, accessToken)).Error().Should(MatchError(jwt.ErrTokenInvalidIssuer))
					},
				)
			},
		)
	},
)
