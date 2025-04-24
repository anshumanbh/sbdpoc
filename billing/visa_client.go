package main

import "os"

// visa_client.go

type VisaClient struct {
	endpoint string
	token    string
}

var visaClient *VisaClient

func initVisaClient() {
	visaClient = &VisaClient{
		endpoint: os.Getenv("VISA_ENDPOINT"),
		token:    os.Getenv("VISA_API_TOKEN"),
	}
}

func (v *VisaClient) Validate(req PaymentRequest) (bool, error) {
	// TODO: call VISA API
	return true, nil
}
