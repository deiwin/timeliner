package i3monitor_test

import (
	. "github.com/deiwin/timeliner/i3monitor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("I3monitor", func() {
	It("should return", func() {
		Expect(SubscribeToActiveWindowUpdates()).NotTo(BeNil())
	})
})
