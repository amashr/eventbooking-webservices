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
			eventsRouter := r.PathPrefix("/events").Subrouter()     
			eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler) 
			eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler) 
			eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

			locationRouter := r.PathPrefix("/locations").Subrouter()
			locationRouter.Methods("GET").Path("").HandlerFunc(handler.allLocationsHandler)
			locationRouter.Methods("POST").Path("").HandlerFunc(handler.newLocationHandler)

			return http.ListenAndServe(endpoint, r)
	} 

	