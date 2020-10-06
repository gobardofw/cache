package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisCache struct {
	prefix string
	pool   *redis.Pool
}

func (c *redisCache) init(prefix string, host string, maxIdle int, maxActive int, db uint8) {
	c.prefix = prefix
	c.pool = &redis.Pool{
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("SELECT", db)
			return conn, err
		},
	}
}

func (c *redisCache) client() redis.Conn {
	return c.pool.Get()
}

func (c *redisCache) prefixer(key string) string {
	if c.prefix == "" {
		return key
	}
	return c.prefix + "-" + key
}

// Put a new value to cache
func (c *redisCache) Put(key string, value interface{}, ttl time.Duration) bool {
	if _, err := c.client().Do("SET", c.prefixer(key), value, "EX", int64(ttl/time.Second)); err == nil {
		return true
	}
	return false
}

// PutForever put value with infinite ttl
func (c *redisCache) PutForever(key string, value interface{}) bool {
	if _, err := c.client().Do("SET", c.prefixer(key), value); err == nil {
		return true
	}
	return false
}

// Set Change value of cache item
func (c *redisCache) Set(key string, value interface{}) bool {
	if _, err := c.client().Do("SET", c.prefixer(key), value, "KEEPTTL"); err == nil {
		return true
	}
	return false
}

// Get item from cache
func (c *redisCache) Get(key string) (interface{}, bool) {
	if value, err := c.client().Do("GET", c.prefixer(key)); err == nil {
		return value, true
	}
	return nil, false
}

// Pull item from cache and remove it
func (c *redisCache) Pull(key string) (interface{}, bool) {
	if value, err := c.client().Do("GET", c.prefixer(key)); err == nil {
		if _, err = c.client().Do("DEL", c.prefixer(key)); err == nil {
			return value, true
		}
	}
	return nil, false
}

// Check if item exists in cache
func (c *redisCache) Exists(key string) bool {
	reply, _ := redis.Int(c.client().Do("EXISTS", c.prefixer(key)))
	return reply == 1
}

// Forget item from cache (delete item)
func (c *redisCache) Forget(key string) bool {
	if _, err := c.client().Do("DEL", c.prefixer(key)); err == nil {
		return true
	}
	return false
}

// TTL get cache item ttl
func (c *redisCache) TTL(key string) (time.Duration, bool) {
	ttl, err := redis.Int(c.client().Do("TTL", c.prefixer(key)))
	if err != nil || ttl == -1 || ttl == -2 {
		return 0, false
	}

	return time.Duration(ttl) * time.Second, true
}

// Bool parse dependency as boolean
func (c *redisCache) Bool(key string, fallback bool) (bool, bool) {
	val, err := redis.Bool(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return val, true
	}
	return fallback, false
}

// Int parse dependency as int
func (c *redisCache) Int(key string, fallback int) (int, bool) {
	val, err := redis.Int64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return int(val), true
	}
	return fallback, false
}

// Int8 parse dependency as int8
func (c *redisCache) Int8(key string, fallback int8) (int8, bool) {
	val, err := redis.Int64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return int8(val), true
	}
	return fallback, false
}

// Int16 parse dependency as int16
func (c *redisCache) Int16(key string, fallback int16) (int16, bool) {
	val, err := redis.Int64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return int16(val), true
	}
	return fallback, false
}

// Int32 parse dependency as int32
func (c *redisCache) Int32(key string, fallback int32) (int32, bool) {
	val, err := redis.Int64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return int32(val), true
	}
	return fallback, false
}

// Int64 parse dependency as int64
func (c *redisCache) Int64(key string, fallback int64) (int64, bool) {
	val, err := redis.Int64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return val, true
	}
	return fallback, false
}

// UInt parse dependency as uint
func (c *redisCache) UInt(key string, fallback uint) (uint, bool) {
	val, err := redis.Uint64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return uint(val), true
	}
	return fallback, false
}

// UInt8 parse dependency as uint8
func (c *redisCache) UInt8(key string, fallback uint8) (uint8, bool) {
	val, err := redis.Uint64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return uint8(val), true
	}
	return fallback, false
}

// UInt16 parse dependency as uint16
func (c *redisCache) UInt16(key string, fallback uint16) (uint16, bool) {
	val, err := redis.Uint64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return uint16(val), true
	}
	return fallback, false
}

// UInt32 parse dependency as uint32
func (c *redisCache) UInt32(key string, fallback uint32) (uint32, bool) {
	val, err := redis.Uint64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return uint32(val), true
	}
	return fallback, false
}

// UInt64 parse dependency as uint64
func (c *redisCache) UInt64(key string, fallback uint64) (uint64, bool) {
	val, err := redis.Uint64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return uint64(val), true
	}
	return fallback, false
}

// Float32 parse dependency as float64
func (c *redisCache) Float32(key string, fallback float32) (float32, bool) {
	val, err := redis.Float64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return float32(val), true
	}
	return fallback, false
}

// Float64 parse dependency as float64
func (c *redisCache) Float64(key string, fallback float64) (float64, bool) {
	val, err := redis.Float64(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return val, true
	}
	return fallback, false
}

// String parse dependency as string
func (c *redisCache) String(key string, fallback string) (string, bool) {
	val, err := redis.String(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return val, true
	}
	return fallback, false
}

// Bytes parse dependency as bytes array
func (c *redisCache) Bytes(key string, fallback []byte) ([]byte, bool) {
	val, err := redis.Bytes(c.client().Do("GET", c.prefixer(key)))
	if err == nil {
		return val, true
	}
	return fallback, false
}

// Increment numeric item in cache
func (c *redisCache) Increment(key string) bool {
	if _, err := c.client().Do("INCR", c.prefixer(key)); err != nil {
		return true
	}
	return false
}

// IncrementBy numeric item in cache by number
func (c *redisCache) IncrementBy(key string, value interface{}) bool {
	var stmt = "INCRBY"
	switch value.(type) {
	case float32, float64:
		stmt = "INCRBYFLOAT"
	}
	if _, err := c.client().Do(stmt, c.prefixer(key), value); err == nil {
		return true
	}
	return false
}

// Decrement numeric item in cache
func (c *redisCache) Decrement(key string) bool {
	if _, err := c.client().Do("DECR", c.prefixer(key)); err == nil {
		return true
	}
	return false
}

// DecrementBy numeric item in cache by number
func (c *redisCache) DecrementBy(key string, value interface{}) bool {
	var stmt = "DECRBY"
	switch value.(type) {
	case float32:
		stmt = "INCRBYFLOAT"
		value = -1 * value.(float32)
	case float64:
		stmt = "INCRBYFLOAT"
		value = -1 * value.(float64)
	}
	if _, err := c.client().Do(stmt, c.prefixer(key), value); err == nil {
		return true
	}
	return false
}
