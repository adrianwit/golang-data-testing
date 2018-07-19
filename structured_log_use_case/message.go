package viant

import (
	"github.com/viant/toolbox/bridge"
	"time"
)


type Message struct {
	Timestamp *time.Time
	EventTypeId int
	IP string
	Request *bridge.HttpRequest
}

func NewMessage(timestamp *time.Time, eventTypeId int, ip string, request *bridge.HttpRequest) *Message{
	return &Message{
		Timestamp:timestamp,
		EventTypeId:eventTypeId,
		IP:ip,
		Request:request,
	}
}

