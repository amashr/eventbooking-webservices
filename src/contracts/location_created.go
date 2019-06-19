
package contracts

import "github.com/amaumba1/eventbooking/src/lib/persistence"

// LocationCreatedEvent is emitted whenever a location is created
type LocationCreatedEvent struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Address      string             `json:"address"`
	Country      string             `json:"country"`
	Halls        []persistence.Hall `json:"hall"`
}

// EventName returns the event's name
func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}