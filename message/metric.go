package message

type Metric struct {
	Message
}

func (e *Metric) Complete() {
	if e.Message.flush != nil {
		e.Message.flush(e)
	}
}

func NewMetric(mType, name string, flush Flush) *Metric {
	return &Metric{
		Message: NewMessage(mType, name, flush),
	}
}
