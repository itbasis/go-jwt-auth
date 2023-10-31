package shared_test

import (
	"context"

	model2 "github.com/itbasis/go-jwt-auth/v2/grpc/shared"
	"github.com/itbasis/go-jwt-auth/v2/model"
	_ "github.com/itbasis/go-test-utils/v2"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
)

type AnotherStruct struct{}

var _ = Describe(
	"SessionUser", func() {
		DescribeTable(
			"Fail", func(testUser any, expectErr error) {
				ctx := context.WithValue(context.Background(), model.SessionUser{}, testUser)
				sessionUser, err := model2.GetSessionUser(ctx)

				Ω(sessionUser).To(BeNil())
				Ω(err).Should(MatchError(ContainSubstring(expectErr.Error())))
			},
			Entry("When SessionUser is nil", nil, model.ErrSessionWithoutAuth),
			Entry("When SessionUser not is interface", &AnotherStruct{}, model.ErrSessionWithoutAuth),
			Entry("When SessionUser is invalid", &model.SessionUser{}, model.ErrSessionInvalidUser),
		)
	},
)
