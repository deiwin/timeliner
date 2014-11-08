package i3monitor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestI3monitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "I3monitor Suite")
}
