package jwttoken_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/gofrs/uuid/v5"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/jwt-token/impl"
	itbasisJwtAuthModel "github.com/itbasis/go-jwt-auth/model"
	itbasisLogUtils "github.com/itbasis/go-log-utils"
	itbasisTestUtils "github.com/itbasis/go-test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
)

var sessionUserTestData = []itbasisJwtAuthModel.SessionUser{
	{Username: "test-user"},
	{UID: uuid.FromStringOrNil("39910003-f693-44ae-979f-f83714a6d459"), Username: "test-user"},
}

func TestJwtToken(t *testing.T) {
	RegisterFailHandler(Fail)

	Ω(sessionUserTestData).NotTo(BeEmpty())

	RunSpecs(t, "JwtToken")
}

var _ = Describe(
	"JwtToken", func() {
		itbasisLogUtils.ConfigureDefaultContextLogger(false)

		BeforeEach(
			func() {
				originalEnvJwtSecretKey := os.Getenv(envJwtSecretKey)
				DeferCleanup(
					func() {
						Ω(os.Setenv(envJwtSecretKey, originalEnvJwtSecretKey)).Should(Succeed())
					},
				)
			},
		)

		for i, testSessionUser := range sessionUserTestData {
			It(
				fmt.Sprintf("Test #%d", i), func() {
					ctx := itbasisTestUtils.TestLoggerWithContext(context.Background())
					mockClock := clock.NewMock()
					mockClock.Set(time.Now())

					Ω(os.Setenv(envJwtSecretKey, "test-key")).Should(Succeed())

					jwtToken, err := itbasisJwtTokenImpl.NewJwtToken(mockClock)
					Ω(err).Should(Succeed())

					accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
					Ω(err).Should(Succeed())
					Ω(accessToken).NotTo(BeEmpty())
					log.Info().Msgf("accessToken: %s", accessToken)

					sessionUser, err := jwtToken.Parse(ctx, accessToken)
					Ω(err).Should(Succeed())
					Ω(*sessionUser).To(Equal(testSessionUser))
				},
			)
		}
	},
)
