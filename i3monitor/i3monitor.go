package i3monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/deiwin/timeliner/throttle"
	"github.com/proxypoke/i3ipc"
)

type ActiveWindowUpdate struct {
	Workspace string
	Window    string
}

const throttleDuration = 1 * time.Second

func SubscribeToActiveWindowUpdates() (out chan ActiveWindowUpdate) {
	out = make(chan ActiveWindowUpdate)

	workspaceEvents, err := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	if err != nil {
		log.Fatal(err)
	}

	throttledWorkspaceEvents := throttle.Throttle(workspaceEvents, throttleDuration)

infninteLoop:
	for {
		select {
		case event := <-throttledWorkspaceEvents:
			fmt.Printf("Received an event: %#v\n", event)
		case <-time.After(5 * time.Second):
			break infninteLoop
		}
	}

	return
}
