package message

type Event struct {
	Message
}

func (e *Event) Complete() {
	if e.Message.flush != nil {
		e.Message.flush(e)
	}
}

func NewEvent(mType, name string, flush Flush) *Event {
	return &Event{
		Message: NewMessage(mType, name, flush),
	}
}
