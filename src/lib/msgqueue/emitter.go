
package msgqueue

// EventEmitter describes an interface for a struct that emits events
type EventEmitter interface {
	Emit(e Event) error 
}