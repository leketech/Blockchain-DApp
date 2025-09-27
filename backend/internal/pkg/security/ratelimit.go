package security

import (
    "sync"
    "time"
)

// RateLimiter holds rate limiting configuration
type RateLimiter struct {
    requests map[string][]time.Time
    mutex    sync.Mutex
    maxRequests int
    window    time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests:    make(map[string][]time.Time),
        maxRequests: maxRequests,
        window:      window,
    }
}

// Allow checks if a request is allowed
func (r *RateLimiter) Allow(key string) bool {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    now := time.Now()
    requests := r.requests[key]

    // Remove old requests outside the window
    validRequests := make([]time.Time, 0)
    for _, reqTime := range requests {
        if now.Sub(reqTime) <= r.window {
            validRequests = append(validRequests, reqTime)
        }
    }

    // Check if we're under the limit
    if len(validRequests) >= r.maxRequests {
        return false
    }

    // Add the new request
    validRequests = append(validRequests, now)
    r.requests[key] = validRequests

    return true
}

// Reset clears the rate limiter for a key
func (r *RateLimiter) Reset(key string) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    delete(r.requests, key)
}