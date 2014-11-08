package throttle_test

import (
	. "github.com/deiwin/timeliner/throttle"
	"github.com/proxypoke/i3ipc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("Throttle", func() {
	var channel chan i3ipc.Event

	BeforeEach(func() {
		channel = make(chan i3ipc.Event)
	})

	It("should return", func() {
		Expect(Throttle(channel, time.Nanosecond)).NotTo(BeNil())
	})

	Context("with a channel throttled by a microsecond", func() {
		var (
			throttled chan i3ipc.Event
			event1    = *new(i3ipc.Event)
			event2    = *new(i3ipc.Event)
			event3    = *new(i3ipc.Event)
			event4    = *new(i3ipc.Event)
		)

		BeforeEach(func() {
			throttled = Throttle(channel, time.Microsecond)
		})

		It("should return the only element pushed", func(done Done) {
			defer close(done)

			channel <- event1
			Expect(<-throttled == event1).To(BeTrue())
		})

		It("should return the last of two consecutively pushed elements", func(done Done) {
			defer close(done)

			channel <- event1
			channel <- event2
			Expect(<-throttled == event2).To(BeTrue())
		})

		It("should return the last of two consecutively pushed elements, twice", func(done Done) {
			defer close(done)

			channel <- event1
			channel <- event2
			Expect(<-throttled == event2).To(BeTrue())

			time.Sleep(2 * time.Microsecond)

			channel <- event3
			channel <- event4
			Expect(<-throttled == event4).To(BeTrue())
		})

		It("should close after input channel is closed", func() {
			close(channel)
			Eventually(throttled).Should(BeClosed())
		})

		It("should be able to receive data and close after input channel is closed", func(done Done) {
			defer close(done)

			channel <- event1
			Expect(<-throttled == event1).To(BeTrue())

			close(channel)
			Eventually(throttled).Should(BeClosed())
		})
	})
})
