package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(rateLimitMiddleware)

	// Proxy payment update to billing service
	r.HandleFunc("/update-payment", proxyHandler)

	addr := ":8080"
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Gateway API listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}
