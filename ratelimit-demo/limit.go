package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var visitors = make(map[string]*visitor)

var mu sync.Mutex

func init() {
	go CleanupVisitors()
}

func getVisitor(ipAddress string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	v, exists := visitors[ipAddress]
	if !exists {
		limiter := rate.NewLimiter(1, 1)
		visitors[ipAddress] = &visitor{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

// CleanupVisitors :
func CleanupVisitors() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		mu.Lock()
		for k, v := range visitors {
			lastseen := v.lastSeen
			if time.Since(lastseen) > 3*time.Minute {
				delete(visitors, k)
			}
		}
		mu.Unlock()
	}
}

func ratelimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("Error while retrieving IpAddress")
			http.Error(w, "InternalServerError", http.StatusInternalServerError)
			return
		}
		ratelimiter := getVisitor(ip)
		if ratelimiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
