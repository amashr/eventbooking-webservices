
package amqp

import (
	"github.com/streadway/amqp"
	//"github.com/amaumba1/eventbooking/src/lib/helper/amqp"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue"
	"github.com/amaumba1/eventbooking/src/contracts"
	"encoding/json"
	"fmt"
	
)

const eventNameHeader = "x-event-name"

type amqpEventListener struct {
	connection *amqp.Connection
	queue string 
	exchange string 
	//mapper msgqueue.EventMapper
}

// func NewAMQPEventListenerFromEnvironment() (msgqueue.EventListener, error) {
// 	var url string
// 	var exchange string
// 	var queue string 

// 	if url = os.Getenv("AMQP_URL"); url == "" {
// 		url = "amqp://localhost:5672"
// 	}
// 	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
// 		exchange = "example"
// 	}

// 	if queue = os.Getenv("AMQP_QUEUE"); queue == "" {
// 		queue = "example"
// 	}

// 	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
// 	return NewAMQPEventListener(conn, exchange, queue)

// }

func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(a.queue, true, false, false, false , nil)
	if err != nil {
		return fmt.Errorf("could not declare queue %s: %s", a.queue, err)
	}
	return nil
}

func NewAMQPEventListener(conn *amqp.Connection, queue string) (msgqueue.EventListener, error) {
	listener := amqpEventListener{
		connection: conn,
		//exchange: exchange,
		queue: queue,
		//mapper: msgqueue.NewEventMapper(),
	}

	err := listener.setup()
	if err != nil {
		return nil, err
	}
	return &listener, nil
}

func (l *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	channel, err := l.connection.Channel()
	if err != nil {
		return nil, nil, err
	}
	defer channel.Close()

	for _, eventName := range eventNames {
		if err := channel.QueueBind(l.queue, eventName, "event", false, nil); err != nil {
			return nil, nil, err
		}
	}
	msgs, err := channel.Consume(l.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}
	events := make(chan msgqueue.Event)
	errors := make(chan error)

	go func() {
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("msg did not contain x-event-name header")
				msg.Nack(false, false)
				continue
			}
			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf(
					"x-event-name header is not string, but %t",
					rawEventName)
					msg.Nack(false, false)
					continue
			}

			var event msgqueue.Event

			switch eventName {
			case "event.created":
				event = new(contracts.EventCreatedEvent)
			default:
				errors <- fmt.Errorf("event type %s is unknown", eventName)
				continue
			}

			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}
			// event, err := l.mapper.MapEvent(eventName, msg.Body)
			// if err != nil {
			// 	errors <- fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
			// 	msg.Nack(false, false)
			// 	continue
			// }
			events <- event
			msg.Ack(false)
		}
	}()
	return events, errors, nil 

}

// func (l *amqpEventListener) Mapper() msgqueue.EventMapper {
// 	return l.mapper
// }