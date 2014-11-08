package throttle

import "time"

type Throttleable interface{}

func Throttle(in <-chan int, duration time.Duration) (out chan int) {
	out = make(chan int)

	go func() {
		var (
			lastMessage    int
			messagePending bool
			ok             bool
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
				case lastMessage, ok = <-in:
					if ok {
						messagePending = true
					} else {
						break
					}
				}
			}
		}
	}()

	return
}
