package main

import (
	"net/http"
)

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid, tenant := validateToken(token)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Store tenant in context
		ctx := r.Context()
		ctx = contextWithTenant(ctx, tenant)
		next(w, r.WithContext(ctx))
	}
}

// stub
func validateToken(tok string) (bool, string) {
	// TODO: implement real token validation
	return true, "tenant-123"
}
