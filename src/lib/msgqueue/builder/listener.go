package builder

import (
	"errors"
	"log"
	"os"

	"github.com/amaumba1/eventbooking/src/lib/msgqueue"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue/kafka"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue/amqp"
)


func NewEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var listener msgqueue.EventListener
	var err error

	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		log.Printf("connecting to kafka brokers at %s", brokers)

		listener, err = kafka.NewKafkaEventEmitterFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else if url := os.Getenv("AMQP_URL"); url != "" {
		log.Printf("connecting to AMQP broker at %s", url)

		listener, err = amqp.NewAMQPEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New("Neither KAFKA_BROKERS nor AMQP_URL specified")
		}
		return listener, nil
	}
