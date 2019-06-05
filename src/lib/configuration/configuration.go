package configuration

import (
		"encoding/json"
		"fmt"
		"os"
		"github.com/amaumba1/eventbooking/src/lib/persistence/dblayer"
)

var ( 
      DBTypeDefault       = dblayer.DBTYPE("mongodb") 
      DBConnectionDefault = "mongodb://127.0.0.1" 
      RestfulEPDefault    = "localhost:8181" 
			MessageBrokerTypeDefault = "amqp"
			AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
    )

type ServiceConfig struct { 
     Databasetype      dblayer.DBTYPE `json:"databasetype"` 
     DBConnection      string         `json:"dbconnection"` 
     RestfulEndpoint   string         `json:"restfulapi_endpoint"` 
		 RestfulTLSEndPint string         `json:"restfulapi-tlsendpoint"`
		 AMQPMessageBroker string   `json:"amap_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) { 
   conf := ServiceConfig{ 
               DBTypeDefault, 
               DBConnectionDefault, 
							 RestfulEPDefault,
							 MessageBrokerTypeDefault,
							 AMQPMessageBrokerDefault,
							}
					
   file, err := os.Open(filename) 
   if err != nil { 
       fmt.Println("Configuration file not found. Continuing with default values.") 
       return conf, err 
    }
	 err = json.NewDecoder(file).Decode(&conf) 
	 if broker :=os.Getenv("AMQP_URL"); broker != "" {
		 conf.AMQPMessageBroker = broker
	 }
   return conf, err
}
