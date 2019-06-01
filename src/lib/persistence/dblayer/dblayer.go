package dblayer

import (
	"github.com/amaumba1/eventbooking/src/lib/persistence"
	"github.com/amaumba1/eventbooking/src/lib/persistence/mongolayer"
)

type DBTYPE string

const (
  MONGODB DBTYPE = "mongodb"
  DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

  switch options {
  case MONGODB:
    return mongolayer.NewMongoDBLayer(connection)
  }
  return nil, nil
}
