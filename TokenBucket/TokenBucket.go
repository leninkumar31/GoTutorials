package main

import (
	"sync"
	"time"
)

// TokenBucket :
type TokenBucket struct {
	Rate      int
	Burst     int
	Available int
	mu        sync.RWMutex
}

// NewTokenBucket :
func NewTokenBucket(r, b int) *TokenBucket {
	tokenBucket := &TokenBucket{
		Rate:      r,
		Burst:     b,
		Available: b,
	}
	go tokenBucket.Fill()
	return tokenBucket
}

// Fill :
func (tokenBucket *TokenBucket) Fill() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		tokenBucket.mu.Lock()
		defer tokenBucket.mu.Unlock()
		tokenBucket.Available = max(tokenBucket.Burst, tokenBucket.Available+tokenBucket.Rate)
	}
}

// Take :
func (tokenBucket *TokenBucket) Take() bool {
	tokenBucket.mu.Lock()
	defer tokenBucket.mu.Unlock()
	if tokenBucket.Available > 0 {
		tokenBucket.Available--
		return true
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
