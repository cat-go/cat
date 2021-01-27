package message

type Heartbeat struct {
	Message
}

func (e *Heartbeat) Complete() {
	if e.Message.flush != nil {
		e.Message.flush(e)
	}
}

func NewHeartbeat(mType, name string, flush Flush) *Heartbeat {
	return &Heartbeat{
		Message: NewMessage(mType, name, flush),
	}
}
