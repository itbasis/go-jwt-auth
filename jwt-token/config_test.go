package jwttoken_test

import (
	"context"
	"os"
	"time"

	"github.com/gofrs/uuid/v5"
	itbasisCoreUtils "github.com/itbasis/go-core-utils/v2"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	envJwtSecretKey                     = "JWT_SECRET_KEY" //nolint:gosec
	envJwtSigningMethod                 = "JWT_SIGNING_METHOD"
	envJwtAccessTokenDurationInSeconds  = "JWT_ACCESS_TOKEN_DURATION"  //nolint:gosec
	envJwtRefreshTokenDurationInSeconds = "JWT_REFRESH_TOKEN_DURATION" //nolint:gosec
)

var _ = Describe(
	"JWT Config", func() {
		var config itbasisJwtToken.Config

		AfterEach(
			func() {
				Ω(os.Unsetenv(envJwtSecretKey)).Should(Succeed())
				Ω(os.Unsetenv(envJwtSigningMethod)).Should(Succeed())
				Ω(os.Unsetenv(envJwtAccessTokenDurationInSeconds)).Should(Succeed())
				Ω(os.Unsetenv(envJwtRefreshTokenDurationInSeconds)).Should(Succeed())

				config = itbasisJwtToken.Config{}
			},
		)

		It(
			"test empty config", func() {
				Ω(itbasisCoreUtils.ReadEnvConfig(context.Background(), &config, nil)).Should(Succeed())
				Ω(config.JwtSecretKey).To(BeEmpty())
				Ω(config.JwtSigningMethod).To(Equal("HS512"))
				Ω(config.JwtAccessTokenDurationInSeconds).To(Equal(itbasisJwtToken.TokenDuration(itbasisJwtToken.DefaultAccessTokenDuration)))
				Ω(config.JwtRefreshTokenDurationInSeconds).To(Equal(itbasisJwtToken.TokenDuration(itbasisJwtToken.DefaultRefreshTokenDuration)))
			},
		)

		It(
			"test custom config", func() {
				testKey, err := uuid.NewV4()
				Ω(err).Should(Succeed())

				Ω(os.Setenv(envJwtSecretKey, testKey.String())).Should(Succeed())
				Ω(os.Setenv(envJwtSigningMethod, "HS256")).Should(Succeed())
				Ω(os.Setenv(envJwtAccessTokenDurationInSeconds, "20")).Should(Succeed())
				Ω(os.Setenv(envJwtRefreshTokenDurationInSeconds, "30")).Should(Succeed())

				Ω(itbasisCoreUtils.ReadEnvConfig(context.Background(), &config, nil)).Should(Succeed())
				Ω(config.JwtSecretKey).To(Equal(testKey.String()))
				Ω(config.JwtSigningMethod).To(Equal("HS256"))
				Ω(config.JwtAccessTokenDurationInSeconds).To(Equal(itbasisJwtToken.TokenDuration(time.Second * 20)))
				Ω(config.JwtRefreshTokenDurationInSeconds).To(Equal(itbasisJwtToken.TokenDuration(time.Second * 30)))
			},
		)
	},
)
