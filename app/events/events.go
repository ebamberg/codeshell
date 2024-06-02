package events

type ApplicationEvent struct {
	Eventtype string
	Payload   any
}

const FS_WORKDIR_CHANGED = "FS_WORKDIR_CHANGED"
const PROFILE_ACTIVATED = "PROFILE_ACTIVATED"
const APPLICATION_ACTIVATED = "PROFILE_ACTIVATED"

type EventListener func(ApplicationEvent)

var eventListeners []EventListener
var eventListenersAsync []EventListener

func RegisterListener(listener EventListener) {
	eventListeners = append(eventListeners, listener)
}

func RegisterAsyncListener(listener EventListener) {
	eventListenersAsync = append(eventListeners, listener)
}

func Broadcast(event ApplicationEvent) {
	for _, listener := range eventListenersAsync {
		go listener(event)
	}
	for _, listener := range eventListeners {
		listener(event)
	}
}
