package throttle

import (
	"time"

	"github.com/proxypoke/i3ipc"
)

// Throttle will filter out messages from a channel. Messages will only appear
// on the resulting channel if they are *NOT* followed by another message within
// the specified amount of time.
func Throttle(in <-chan i3ipc.Event, duration time.Duration) (out chan i3ipc.Event) {
	out = make(chan i3ipc.Event)
	go throttleSynchronously(in, out, duration)
	return
}

func throttleSynchronously(in <-chan i3ipc.Event, out chan i3ipc.Event, duration time.Duration) {
	var (
		lastMessage    i3ipc.Event
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
