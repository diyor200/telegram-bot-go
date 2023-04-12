package events

type Events interface {
	Fetch(linmit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type    Type
	Message string
}
