package rest

import (
	"net/http"

	"github.com/amaumba1/eventbooking/src/lib/persistence"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue"

	"github.com/gorilla/mux"
)

	func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error { 
			handler := NewEventHandler(dbHandler, eventEmitter)
			r := mux.NewRouter() 
			eventsrouter := r.PathPrefix("/events").Subrouter()     
			eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler) 
			eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler) 
			eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

			return http.ListenAndServe(endpoint, r)
	} 

	