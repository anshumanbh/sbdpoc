package main

import (
	"net/http"
	"sync"
	"time"
)

// Simple per-user rate limiter
var (
	mu        sync.Mutex
	userCalls = make(map[string][]time.Time)
)

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-Workspace-Owner")
		now := time.Now()

		mu.Lock()
		calls := userCalls[user]
		// Filter calls in last minute
		var recent []time.Time
		for _, t := range calls {
			if now.Sub(t) < time.Minute {
				recent = append(recent, t)
			}
		}
		if len(recent) >= 100 {
			mu.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		userCalls[user] = append(recent, now)
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
