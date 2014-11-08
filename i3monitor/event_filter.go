package i3monitor

import "github.com/samuelotter/i3ipc"

func FilterTitleAndFocusEvents(in <-chan i3ipc.Event) chan i3ipc.Event {
	return filter(in, func(event i3ipc.Event) bool {
		return event.Change == "title" || event.Change == "focus"
	})
}

func filter(in <-chan i3ipc.Event, predicate func(i3ipc.Event) bool) (out chan i3ipc.Event) {
	out = make(chan i3ipc.Event)

	go func() {
		for event := range in {
			if predicate(event) {
				out <- event
			}
		}
		close(out)
	}()

	return
}
