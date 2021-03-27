package dblayer

import (
	"github.com/reserva/v1/myevents/src/lib/presistence"
	"github.com/reserva/v1/myevents/src/lib/presistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	DYNAMOD DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (presistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
