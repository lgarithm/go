package ratelimit

import (
	"sync"
	"time"
)

// RateLimitor can block a token for given duration
type RateLimitor struct {
	sync.Mutex
	blocked map[string]bool
}

// New creates a new RateLimitor
func New() *RateLimitor {
	r := RateLimitor{
		blocked: make(map[string]bool),
	}
	return &r
}

// BlockAll blocks a list of keys for their durations
func (r *RateLimitor) BlockAll(m map[string]time.Duration) bool {
	r.Lock()
	defer r.Unlock()
	for key := range m {
		if _, ok := r.blocked[key]; ok {
			return false
		}
	}
	for key, d := range m {
		r.blocked[key] = true
		go func(key string, d time.Duration) {
			<-time.After(d)
			r.unblock(key)
		}(key, d)
	}
	return true
}

func (r *RateLimitor) unblock(key string) {
	r.Lock()
	defer r.Unlock()
	delete(r.blocked, key)
}
