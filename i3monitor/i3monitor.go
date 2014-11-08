package i3monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/deiwin/timeliner/throttle"
	"github.com/samuelotter/i3ipc"
)

type ActiveWindowUpdate struct {
	application string
	title       string
}

const throttleDuration = 1 * time.Second

func SubscribeToActiveWindowUpdates() (out chan ActiveWindowUpdate) {
	out = make(chan ActiveWindowUpdate)

	windowEvents, err := i3ipc.Subscribe(i3ipc.I3WindowEvent)
	if err != nil {
		log.Fatal(err)
	}

	throttledWindowEvents := throttle.Throttle(windowEvents, throttleDuration)

infninteLoop:
	for {
		select {
		case event := <-throttledWindowEvents:
			fmt.Printf("Received an event: %#v\n", event)
		case <-time.After(5 * time.Second):
			break infninteLoop
		}
	}

	return
}
