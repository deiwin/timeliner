package i3monitor_test

import (
	. "github.com/deiwin/timeliner/i3monitor"
	"github.com/samuelotter/i3ipc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event Filter", func() {
	var channel chan i3ipc.Event

	BeforeEach(func() {
		channel = make(chan i3ipc.Event)
	})

	It("should return", func() {
		Expect(FilterTitleAndFocusEvents(channel)).NotTo(BeNil())
	})

	Describe("filtered channel", func() {
		var filteredChannel chan i3ipc.Event

		BeforeEach(func() {
			filteredChannel = FilterTitleAndFocusEvents(channel)
		})

		Context("with random event", func() {
			var randomEvent = i3ipc.Event{Change: "random"}

			It("should ignore a random event", func(done Done) {
				defer close(done)

				channel <- randomEvent
				Consistently(filteredChannel).ShouldNot(Receive())
			})

			Context("with focus and title events", func() {
				var focusEvent = i3ipc.Event{Change: "focus"}
				var titleEvent = i3ipc.Event{Change: "title"}

				It("should forward the focus event to the output channel", func(done Done) {
					defer close(done)

					channel <- focusEvent
					Expect(<-filteredChannel).To(Equal(focusEvent))
				})

				It("should forward the title event to the output channel", func(done Done) {
					defer close(done)

					channel <- titleEvent
					Expect(<-filteredChannel).To(Equal(titleEvent))
				})

				It("should ignore random events between title and focus events", func(done Done) {
					defer close(done)

					channel <- titleEvent
					Expect(<-filteredChannel).To(Equal(titleEvent))
					channel <- randomEvent
					channel <- randomEvent
					channel <- focusEvent
					Expect(<-filteredChannel).To(Equal(focusEvent))
					channel <- randomEvent
					channel <- randomEvent
					close(channel)
					Eventually(filteredChannel).Should(BeClosed())
				})
			})
		})

		It("should close when input channel closed", func() {
			close(channel)
			Eventually(filteredChannel).Should(BeClosed())
		})
	})

})
