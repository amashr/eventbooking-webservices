package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/amaumba1/eventbooking/src/lib/persistence"
	"github.com/amaumba1/eventbooking/src/lib/msgqueue"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	srv.ListenAndServe()
}
