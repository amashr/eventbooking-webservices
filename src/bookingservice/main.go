
package main

import (
	"github.com/streadway/amqp"
	"github.com/amaumba1/eventbooking/src/lib/configuration"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue/amqp"
	"flag"
)

func main() {
	confPath := flag.String("config","./configuration/config.json","path to config file")
	flag.Parse()
	config := configuration.ExtractConfiguration(*confPath)

	dblayer, err := dblayer.NewPersistenceLayer(config.Databsetype, config.DBConnection)
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	eventListener, err := msgqueue_amqp.NewAMQPEventListener(conn)
	if err != nil {
		panic(err)
	}
	processor := &listener.EventProcessor{eventListener, dblayer}
	go processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, dblayer, eventEmitter)
}