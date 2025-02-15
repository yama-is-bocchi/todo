package ui

import "github.com/nsf/termbox-go"

type EventHandlerInterface interface {
	HandleEvent() termbox.Event
}

type eventHandler struct{}

func NewEventHandler() *eventHandler {
	return &eventHandler{}
}

func (eventHandler eventHandler) HandleEvent() termbox.Event {
	return termbox.PollEvent()
}
