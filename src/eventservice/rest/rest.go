package rest

import (
	"net/http"

	"github.com/amaumba1/eventbooking/src/lib/persistence"

	"github.com/gorilla/mux"
)


func serveAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databasehandler)
	
	//Create a router using the gorilla mux package 
	r := mux.NewRouter()
	//A subrouter method which createt a new object that will be in charge of any incoming HTTP request directed towards a relative URL that start with /events. A new router assign will be called eventsrouter 
	eventsrouter := r.PathPrefix("/events").Subrouter() // PathPrefix method is used to capture any URL path that starts with /events. 
	
	// A new router eventsrouter object will be used to define what to do with the rest of the URLs that share the /events prefix 
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}