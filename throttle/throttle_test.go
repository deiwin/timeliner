package throttle_test

import (
	. "github.com/deiwin/timeliner/throttle"
	"github.com/samuelotter/i3ipc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("Throttle", func() {
	var channel chan i3ipc.Event

	BeforeEach(func() {
		channel = make(chan i3ipc.Event)
	})

	It("should return", func(done Done) {
		defer close(done)

		Expect(Throttle(channel, time.Nanosecond)).NotTo(BeNil())
	})

	Context("with a channel throttled by a microsecond", func() {
		var (
			throttled chan i3ipc.Event
			event1    = i3ipc.Event{Change: "1"}
			event2    = i3ipc.Event{Change: "2"}
			event3    = i3ipc.Event{Change: "3"}
			event4    = i3ipc.Event{Change: "4"}
		)

		BeforeEach(func() {
			throttled = Throttle(channel, time.Microsecond)
		})

		It("should return the only element pushed", func(done Done) {
			defer close(done)

			channel <- event1
			Expect(<-throttled).To(Equal(event1))
		})

		It("should return the last of two consecutively pushed elements", func(done Done) {
			defer close(done)

			channel <- event1
			channel <- event2
			Expect(<-throttled).To(Equal(event2))
		})

		It("should return the last of two consecutively pushed elements, twice", func(done Done) {
			defer close(done)

			channel <- event1
			channel <- event2
			Expect(<-throttled).To(Equal(event2))

			time.Sleep(2 * time.Microsecond)

			channel <- event3
			channel <- event4
			Expect(<-throttled).To(Equal(event4))
		})

		It("should close after input channel is closed", func(done Done) {
			defer close(done)

			close(channel)
			Eventually(throttled).Should(BeClosed())
		})

		It("should be able to receive data and close after input channel is closed", func(done Done) {
			defer close(done)

			channel <- event1
			Expect(<-throttled).To(Equal(event1))

			close(channel)
			Eventually(throttled).Should(BeClosed())
		})

		It("close should also work right after put", func(done Done) {
			defer close(done)

			channel <- event1
			close(channel)
			Consistently(throttled).ShouldNot(Receive())
		})
	})
})
