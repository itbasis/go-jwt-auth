package shared_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
)

func TestGetSessionUser(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "SessionUser")
}
