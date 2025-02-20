package ui

import (
	"github.com/nsf/termbox-go"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (eventHandler EventHandler) HandleEvent() termbox.Event {
	return termbox.PollEvent()
}
