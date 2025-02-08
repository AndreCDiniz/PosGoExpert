package events

import "time"

type EventInterface interface {
	// Get the event name
	GetName() string
	// Get the event date time
	GetDateTime() time.Time
	// Get the event payload that can be anything for this reason we use interface{}
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	// Handle an event
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	// Register an event with a handler
	Register(eventname string, handler EventHandlerInterface) error
	// Dispatch an event
	Dispatch(event EventInterface) error
	// Remove an event handler
	Remove(eventName string, handler EventHandlerInterface) error
	// Check if an event has a handler
	Has(eventName string, handler EventHandlerInterface) bool
	// Clear all events
	Clear() error
}
