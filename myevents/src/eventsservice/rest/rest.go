package rest

import (
	"net/http"

	"github.com/reserva/v1/myevents/src/lib/presistence"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databasehandler presistence.DatabaseHandler) error {
	handler := NewEvenHandler(databasehandler)
	r := mux.NewRouter()

	eventsrouter := r.PathPrefix("/events").Subrouter()

	eventsrouter.Methods("GET").Path("/{SreachCriteria}/{search}").HandlerFunc(handler.FindEventHandler)

	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)

	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
