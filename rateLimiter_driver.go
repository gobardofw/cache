package cache

import (
	"time"
)

// rateLimiterDriver rate limiter driver
type rateLimiterDriver struct {
	Key   string
	Max   uint32
	Cache Cache
}

func (limiter *rateLimiterDriver) init(key string, maxAttempts uint32, ttl time.Duration, cache Cache) {
	limiter.Key = key
	limiter.Max = maxAttempts
	limiter.Cache = cache
	if !cache.Exists(key) {
		cache.Put(key, maxAttempts, ttl)
	}
}

// Hit decrease the allowed times
func (limiter *rateLimiterDriver) Hit() {
	old, _ := limiter.Cache.Int(limiter.Key, 0)
	if old > 0 {
		limiter.Cache.Decrement(limiter.Key)
	}
}

// Lock lock rate limiter
func (limiter *rateLimiterDriver) Lock() {
	if limiter.Cache.Exists(limiter.Key) {
		limiter.Cache.Set(limiter.Key, 0)
	}
}

// Reset reset rate limiter
func (limiter *rateLimiterDriver) Reset() {
	limiter.Cache.Forget(limiter.Key)
}

// MustLock check if rate limiter must lock access
func (limiter *rateLimiterDriver) MustLock() bool {
	old, _ := limiter.Cache.Int(limiter.Key, 0)
	return limiter.Cache.Exists(limiter.Key) && old <= 0
}

// TotalAttempts get user attempts count
func (limiter *rateLimiterDriver) TotalAttempts() uint32 {
	old, _ := limiter.Cache.Int(limiter.Key, 0)
	if old < 0 {
		old = 0
	}
	if limiter.Cache.Exists(limiter.Key) {
		return limiter.Max - uint32(old)
	}
	return 0
}

// RetriesLeft get user retries left
func (limiter *rateLimiterDriver) RetriesLeft() uint32 {
	old, _ := limiter.Cache.Int(limiter.Key, 0)
	if old < 0 {
		old = 0
	}
	if limiter.Cache.Exists(limiter.Key) {
		return limiter.Max - limiter.TotalAttempts()
	}
	return 0
}

// AvailableIn get time until unlock
func (limiter *rateLimiterDriver) AvailableIn() time.Duration {
	ttl, ok := limiter.Cache.TTL(limiter.Key)
	if !ok {
		return 0
	}
	return ttl
}
