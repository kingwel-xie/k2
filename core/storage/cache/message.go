package cache

import (
	"github.com/robinjoseph08/redisqueue/v2"
)

type Message struct {
	redisqueue.Message
}

func (m *Message) GetID() string {
	return m.ID
}

func (m *Message) GetStream() string {
	return m.Stream
}

func (m *Message) GetValues() map[string]interface{} {
	return m.Values
}

func (m *Message) SetID(id string) {
	m.ID = id
}

func (m *Message) SetStream(stream string) {
	m.Stream = stream
}

func (m *Message) SetValues(values map[string]interface{}) {
	m.Values = values
}
