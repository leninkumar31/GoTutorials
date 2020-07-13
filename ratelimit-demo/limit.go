package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var ratelimiter = rate.NewLimiter(1, 1)

func ratelimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ratelimiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
