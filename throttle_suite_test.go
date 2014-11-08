package throttle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoThrottle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Throttle Suite")
}
