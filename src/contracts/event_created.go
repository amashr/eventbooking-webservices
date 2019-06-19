
package contracts

import "time"

type EventCreatedEvent struct {
	ID            string `json:"id"`
	Name          string `json:"nameId"`
	LocationID    string `json:"locationId"`
	Start         time.Time `json:"start_time"`
	End           time.Time `json:"end_time"`
}

func (e *EventCreatedEvent) EventName() string {
	return "event.created"
}

func (e *EventCreatedEvent) PartitionKey() string {
	return e.ID
}