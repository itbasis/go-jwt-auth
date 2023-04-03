package shared_test

import (
	"context"
	"testing"

	"github.com/itbasis/go-jwt-auth/grpc/shared"
	"github.com/itbasis/go-jwt-auth/model"
	itbasisTestUtils "github.com/itbasis/go-test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

type AnotherStruct struct{}

func TestGetSessionUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SessionUser")
}

var _ = Describe(
	"SessionUser", func() {

		DescribeTable(
			"Fail", func(testUser any, expectErr *status.Status) {
				ctx := itbasisTestUtils.TestLoggerWithContext(
					context.WithValue(
						context.Background(),
						model.SessionUser{},
						testUser,
					),
				)
				sessionUser, statusErr := shared.GetSessionUser(ctx)
				Ω(sessionUser).To(BeNil())
				Ω(statusErr).Should(Equal(expectErr))
			},
			Entry("When SessionUser is nil", nil, shared.ErrSessionWithoutAuth),
			Entry("When SessionUser not is interface", &AnotherStruct{}, shared.ErrSessionWithoutAuth),
			Entry("When SessionUser is invalid", &model.SessionUser{}, shared.ErrSessionInvalidUser),
		)
	},
)
