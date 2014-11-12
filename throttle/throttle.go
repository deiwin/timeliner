package throttle

import (
	"time"

	"github.com/samuelotter/i3ipc"
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
	var lastMessage *i3ipc.Event
	defer close(out)

	hasAPendingMessage := func() bool {
		return lastMessage == nil
	}

	for {
		if hasAPendingMessage() {
			select {
			case message, ok := <-in:
				if ok {
					lastMessage = &message
				} else {
					return
				}
			}
		} else {
			select {
			case message, ok := <-in:
				if ok {
					lastMessage = &message
				} else {
					return
				}
			case <-time.After(duration):
				out <- *lastMessage
				lastMessage = nil
			}
		}
	}
}
