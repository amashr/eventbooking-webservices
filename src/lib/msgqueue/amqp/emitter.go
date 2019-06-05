
package amqp

import (
	
	"github.com/streadway/amqp"
	"encoding/json"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
}

func NewAMQPEventEmitter(conn *amqp.Connection) (msgqueue.EventEmitter, error ){
	emitter := &amqpEventEmitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err 
	}
	return emitter, nil
}

func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	msg := amqp.Publishing{
		Headers:    amqp.Table{"x-event-name": event.EventName()},
		Body:       jsonDoc,
		ContentType: "application/json",
	}
	return channel.Publish(
		"event",
		event.EventName(),
		false,
		false, msg)
}

func (a *amqpEventEmitter)setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err 
	}

	defer channel.Close()

	return channel.ExchangeDeclare("event", "topic", true, false, false, false, nil)
}