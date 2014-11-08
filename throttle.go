package throttle

import "time"

type Throttleable interface{}

func Throttle(in <-chan int, duration time.Duration) (out chan int) {
	out = make(chan int)

	go func() {
		var (
			lastMessage    int
			messagePending bool
		)

		for {
			if messagePending {
				select {
				case lastMessage = <-in:
					messagePending = true
				case <-time.After(duration):
					out <- lastMessage
					messagePending = false
				}
			} else {
				select {
				case lastMessage = <-in:
					messagePending = true
				}
				// could also check every once in a while if the in channel is closed to break out
			}
		}
	}()

	return
}
