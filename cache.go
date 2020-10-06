package cache

import "time"

// Cache interface for cache drivers.
type Cache interface {
	// Put a new value to cache
	Put(key string, value interface{}, ttl time.Duration) bool
	// PutForever put value with infinite ttl
	PutForever(key string, value interface{}) bool
	// Set Change value of cache item
	Set(key string, value interface{}) bool
	// Get item from cache
	Get(key string) (interface{}, bool)
	// Pull item from cache and remove it
	Pull(key string) (interface{}, bool)
	// Check if item exists in cache
	Exists(key string) bool
	// Forget item from cache (delete item)
	Forget(key string) bool
	// TTL get cache item ttl
	TTL(key string) (time.Duration, bool)
	// Bool parse dependency as boolean
	Bool(key string, fallback bool) (bool, bool)
	// Int parse dependency as int
	Int(key string, fallback int) (int, bool)
	// Int8 parse dependency as int8
	Int8(key string, fallback int8) (int8, bool)
	// Int16 parse dependency as int16
	Int16(key string, fallback int16) (int16, bool)
	// Int32 parse dependency as int32
	Int32(key string, fallback int32) (int32, bool)
	// Int64 parse dependency as int64
	Int64(key string, fallback int64) (int64, bool)
	// UInt parse dependency as uint
	UInt(key string, fallback uint) (uint, bool)
	// UInt8 parse dependency as uint8
	UInt8(key string, fallback uint8) (uint8, bool)
	// UInt16 parse dependency as uint16
	UInt16(key string, fallback uint16) (uint16, bool)
	// UInt32 parse dependency as uint32
	UInt32(key string, fallback uint32) (uint32, bool)
	// UInt64 parse dependency as uint64
	UInt64(key string, fallback uint64) (uint64, bool)
	// Float32 parse dependency as float64
	Float32(key string, fallback float32) (float32, bool)
	// Float64 parse dependency as float64
	Float64(key string, fallback float64) (float64, bool)
	// String parse dependency as string
	String(key string, fallback string) (string, bool)
	// Bytes parse dependency as bytes array
	Bytes(key string, fallback []byte) ([]byte, bool)
	// Increment numeric item in cache
	Increment(key string) bool
	// IncrementBy numeric item in cache by number
	IncrementBy(key string, value interface{}) bool
	// Decrement numeric item in cache
	Decrement(key string) bool
	// DecrementBy numeric item in cache by number
	DecrementBy(key string, value interface{}) bool
}
