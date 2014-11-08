package throttle

import "time"

// Throttle will filter out messages from a channel. Messages will only appear
// on the resulting channel if they are *NOT* followed by another message within
// the specified amount of time.
func Throttle(in <-chan int, duration time.Duration) (out chan int) {
	out = make(chan int)
	go throttleSynchronously(in, out, duration)
	return
}

func throttleSynchronously(in <-chan int, out chan int, duration time.Duration) {
	var (
		lastMessage    int
		messagePending bool
		ok             bool
	)

infiniteLoop:
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
					close(out)
					break infiniteLoop
				}
			}
		}
	}
}
