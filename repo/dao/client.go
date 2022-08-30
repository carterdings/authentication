package dao

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Client client interface
type Client interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{}, d time.Duration)
	Delete(key string)
}

// New new client
func New(defaultExpiration, cleanupInterval time.Duration) Client {
	return cache.New(defaultExpiration, cleanupInterval)
}

var (
	DefaultClient = New(5*time.Minute, 10*time.Minute)
)

// Get get
func Get(key string) (interface{}, bool) {
	return DefaultClient.Get(key)
}

// Set set
func Set(key string, val interface{}, d time.Duration) {
	DefaultClient.Set(key, val, d)
}

// Delete delete
func Delete(key string) {
	DefaultClient.Delete(key)
}
