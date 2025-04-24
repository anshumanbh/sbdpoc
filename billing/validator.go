package main

import (
	"regexp"
)

var (
	ccNumRegex = regexp.MustCompile(`^[0-9]{13,19}$`)
	expRegex   = regexp.MustCompile(`^(0[1-9]|1[0-2])/[0-9]{2}$`)
	cvvRegex   = regexp.MustCompile(`^[0-9]{3,4}$`)
)

func validateInput(r *PaymentRequest) error {
	if r.OrganizationName == "" {
		return ErrBadRequest
	}
	if !ccNumRegex.MatchString(r.CardNumber) {
		return ErrBadRequest
	}
	if !expRegex.MatchString(r.Expiry) {
		return ErrBadRequest
	}
	if !cvvRegex.MatchString(r.CVV) {
		return ErrBadRequest
	}
	return nil
}
