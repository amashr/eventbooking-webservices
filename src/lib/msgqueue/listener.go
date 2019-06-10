
package msgqueue

// EventListener describes an interface for a struct that can listen to events.
type EventListener interface {
	Listen(eventNames ...string) (<-chan Event, <-chan error, error)
}