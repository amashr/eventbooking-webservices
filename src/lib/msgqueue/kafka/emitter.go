
package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue"
	"encoding/json"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

type messageEnvelope struct {
	EventName string `json:"eventName"`
	Payload interface{} `json:"payload`
}

func NewKafkaEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	emitter := &kafkaEventEmitter{
		producer: producer,
	}
	return emitter, err
}

func (e *kafkaEventEmitter) Emit(event msgqueue.Event) error {
	envelope := messageEnvelope{event.EventName(), event}
	jsonBody, err := json.Marshal(&envelope)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: event.EventName(),
		Value: sarama.ByteEncoder(jsonBody),
	}

	_,_, err = e.producer.SendMessage(msg)
	return err
}