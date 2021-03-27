package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/reserva/v1/myevents/src/lib/presistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBconnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
)

type ServiceConfig struct {
	Databasetype   dblayer.DBTYPE `json:"databasetype"`
	DBConnection   string         `json:"dbconnection"`
	ResfulEndpoint string         `json:"restfulapi_endpoint"`
}

func ExtraConfiguration(filename string) (ServiceConfig, error) {
	DBConnectionDefault := 
	conf := ServiceConfig{
		    DBTypeDefault,
		    DBConnectionDefault,
		    RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
