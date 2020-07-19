package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter = rate.NewLimiter(1, 1)

func ratelimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Limit on Number requests has Exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
