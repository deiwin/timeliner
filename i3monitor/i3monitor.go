package i3monitor

// import "github.com/proxypoke/i3ipc"

type ActiveWindowUpdate struct {
	Workspace string
	Window    string
}

func SubscribeToActiveWindowUpdates() (subs chan ActiveWindowUpdate) {
	subs = make(chan ActiveWindowUpdate)

	return
}
