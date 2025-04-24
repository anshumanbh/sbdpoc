package main

import (
	"io"
	"net/http"
	"os"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Forward to billing service
	billingURL := os.Getenv("BILLING_ENDPOINT") // e.g. http://localhost:8081
	req, err := http.NewRequest(r.Method, billingURL+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
