package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ErrBadRequest      = errors.New("bad input")
	ErrVisaUnavailable = errors.New("visa service unavailable")
)

type PaymentRequest struct {
	TenantID         string `json:"tenant_id"`
	OrganizationName string `json:"organization_name"`
	CardNumber       string `json:"card_number"`
	Expiry           string `json:"expiry"`
	CVV              string `json:"cvv"`
}

func updatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req PaymentRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Inject tenant from context
	req.TenantID = tenantFromContext(r.Context())

	if err := validateInput(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := visaClient.Validate(req)
	if err != nil {
		logError(err, req)
		http.Error(w, "Visa validation failed", http.StatusServiceUnavailable)
		return
	}
	if !ok {
		http.Error(w, "Invalid credit card", http.StatusBadRequest)
		return
	}

	if err := storeCard(req); err != nil {
		logError(err, req)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logInfo(req)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
