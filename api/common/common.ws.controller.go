package common

import (
	"github.com/nsocrates/go-blog/websockets"
)

type WSAction struct {
	Verb    string      `json:"verb"`
	Channel string      `json:"channel"`
	Path    string      `json:"path"`
	Data    interface{} `json:"data"`
}

func NewWSAction() WSAction {
	return WSAction{}
}

func Broadcast(verb, channel, path string, data interface{}) {
	action := NewWSAction()
	action.Verb = verb
	action.Channel = channel
	action.Path = path
	action.Data = data

	websockets.HubInstance.Broadcast() <- &websockets.Message{Channel: channel, Data: action}
}
