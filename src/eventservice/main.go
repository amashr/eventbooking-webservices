package main

import (
	"flag"
	"log"

	"github.com/amaumba1/eventbooking/src/eventservice/rest"
	"github.com/amaumba1/eventbooking/src/lib/configuration"
	"github.com/amaumba1/eventbooking/src/lib/persistence/dblayer"	
	//"github.com/streadway/amqp"
	//"os"
)

func main() {
	confPath := flag.String("conf", `./configuration/config.json`, "flag to set the path to the configuration json file")
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful...")
	//RESTful API Start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint,dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}


