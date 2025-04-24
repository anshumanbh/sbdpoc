package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize dependencies
	initDB()
	initVisaClient()
	initLogger()

	r := mux.NewRouter()
	r.HandleFunc("/update-payment", authenticate(updatePaymentHandler)).Methods("POST")

	addr := ":8081"
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Billing Service listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}
