package throttle_test

import (
	. "github.com/deiwin/timeliner/throttle"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("Throttle", func() {
	var channel chan int

	BeforeEach(func() {
		channel = make(chan int)
	})

	It("should return", func() {
		Expect(Throttle(channel, time.Nanosecond)).NotTo(BeNil())
	})

	Context("with a channel throttled by a microsecond", func() {
		var throttled chan int

		BeforeEach(func() {
			throttled = Throttle(channel, time.Microsecond)
		})

		It("should return the only element pushed", func(done Done) {
			defer close(done)

			channel <- 1
			Expect(<-throttled).To(Equal(1))
		})

		It("should return the last of two consecutively pushed elements", func(done Done) {
			defer close(done)

			channel <- 1
			channel <- 2
			Expect(<-throttled).To(Equal(2))
		})

		It("should return the last of two consecutively pushed elements, twice", func(done Done) {
			defer close(done)

			channel <- 1
			channel <- 2
			Expect(<-throttled).To(Equal(2))

			time.Sleep(2 * time.Microsecond)

			channel <- 3
			channel <- 4
			Expect(<-throttled).To(Equal(4))
		})

		It("should close after input channel is closed", func() {
			close(channel)
			Eventually(throttled).Should(BeClosed())
		})

		It("should be able to receive data and close after input channel is closed", func(done Done) {
			defer close(done)

			channel <- 5
			Expect(<-throttled).To(Equal(5))

			close(channel)
			Eventually(throttled).Should(BeClosed())
		})
	})
})
