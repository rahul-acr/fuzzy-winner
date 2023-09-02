package events

type Event struct {
	Name    string
	Payload interface{}
}

type listeners []Listener

type Listener func(event Event)

var listenersMap = make(map[string]listeners)

func Publish(eventName string, payload interface{}) {
	for _, listener := range listenersMap[eventName] {
		listener(Event{Name: eventName, Payload: payload})
	}
}

func Listen(eventName string, listener Listener) {
	listenersMap[eventName] = append(listenersMap[eventName], listener)
}


func Clear(eventName string) {
	listenersMap[eventName] = make([]Listener, 0)
}