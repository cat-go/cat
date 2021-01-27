package message

import (
	"bytes"
	"time"
)

type NullMessage struct {
	rootMessageId   string
	parentMessageId string
	messageId       string
}

type NullTransaction struct {
	NullMessage
}

var nullMessage = &NullMessage{}

func (m *NullMessage) Complete() {
}

func (m *NullMessage) GetType() string {
	return ""
}

func (m *NullMessage) GetName() string {
	return ""
}

func (m *NullMessage) GetStatus() string {
	return ""
}

func (m *NullMessage) GetData() *bytes.Buffer {
	return nil
}

func (m *NullMessage) GetTime() time.Time {
	return time.Now()
}

func (m *NullMessage) SetTime(time time.Time) {
}

func (m *NullMessage) SetTimestamp(timestampMs int64) {
}

func (m *NullMessage) AddData(k string, v ...string) {
}

func (m *NullMessage) SetData(v string) {
}

func (m *NullMessage) SetStatus(status string) {
}

func (m *NullMessage) GetRootMessageId() string {
	return m.rootMessageId
}

func (m *NullMessage) GetParentMessageId() string {
	return m.parentMessageId
}

func (m *NullMessage) GetMessageId() string {
	return m.messageId
}


func (t *NullTransaction) AddChild(m Messager) {
	return
}

func (t *NullTransaction) GetChildren() []Messager {
	return []Messager{}
}

func (t *NullTransaction) GetDuration() time.Duration {
	return 0
}

func (t *NullTransaction) SetDuration(duration time.Duration) {
}

func (t *NullTransaction) SetDurationStart(duration time.Time) {
}

func (t *NullTransaction) NewEvent(mtype, name string) Messager {
	return nullMessage
}

func (t *NullTransaction) LogEvent(mtype, name string, args ...string) {
	return
}

func (t *NullTransaction) SetRootMessageId(rootMessageId string) {
	t.rootMessageId = rootMessageId
}

func (t *NullTransaction) SetParentMessageId(parentMessageId string) {
	t.parentMessageId = parentMessageId
}

func (t *NullTransaction) SetMessageId(messageId string) {
	t.messageId = messageId
}