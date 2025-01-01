package enums

type Event int

// Event States
const (
	CREATED            = iota
	DELETED            = iota
	MODIFIED           = iota
	EVENT_INITIAL_SCAN = iota
)

var eventNames = map[Event]string{
	CREATED:            "CREATED",
	DELETED:            "DELETED",
	MODIFIED:           "MODIFIED",
	EVENT_INITIAL_SCAN: "INITIAL_SCAN",
}

func (event Event) String() string {
	return eventNames[event]
}
