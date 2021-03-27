package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/reserva/v1/myevents/src/eventservice/rest"
	"github.com/reserva/v1/myevents/src/lib/configuration"
	"github.com/reserva/v1/myevents/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flage to set the path to the configuration json file")
	flag.Parse()
	//extract config
	config, _ := configuration.ExtraConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API starts here
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
