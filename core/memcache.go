package core

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	name string
	mc *memcache.Client
}

func NewMemcache(name string) *Memcache {
	mc := App().GetMemcache().Mc(name)
	return &Memcache{
		name: name,
		mc:   mc,
	}
}

// Get gets the item for the given key.
// The key must be at most 250 bytes in length.
func (m *Memcache) Get(key string) (string, error) {
	item, err := m.mc.Get(key)
	return string(item.Value), err
}

// GetMulti is a batch version of Get. The returned map from keys to
// items may have fewer elements than the input slice, due to memcache
// cache misses. Each key must be at most 250 bytes in length.
// If no error is returned, the returned map will also be non-nil.
func (m *Memcache) GetMulti(keys []string) (map[string]string, error) {
	items, err := m.mc.GetMulti(keys)
	if err != nil {
		return nil, err
	}
	val := map[string]string{}
	for key, item := range items {
		val[key] = string(item.Value)
	}
	return val, nil
}

// Increment atomically increments key by delta. The return value is
// the new value after being incremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On 64-bit overflow, the new value wraps around.
func (m *Memcache) Increment(key string, delta uint64) (newValue uint64, err error) {
	return m.mc.Increment(key, delta)
}

// Decrement atomically decrements key by delta. The return value is
// the new value after being decremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On underflow, the new value is capped at zero and does not wrap
// around.
func (m *Memcache) Decrement(key string, delta uint64) (newValue uint64, err error) {
	return m.mc.Decrement(key, delta)
}


// Set writes the given item, unconditionally.
func (m *Memcache) Set(key string, value string, expiration int32) error {
	item := &memcache.Item{
		Key: key,
		Value: []byte(value),
		Expiration: expiration,
	}
	return m.mc.Set(item)
}

// Add writes the given item, if no value already exists for its
// key. ErrNotStored is returned if that condition is not met.
func (m *Memcache) Add(key string, value string, expiration int32) error {
	item := &memcache.Item{
		Key: key,
		Value: []byte(value),
		Expiration: expiration,
	}
	return m.mc.Add(item)
}

// Replace writes the given item, but only if the server *does*
// already hold data for this key
func (m *Memcache) Replace(key string, value string, expiration int32) error {
	item := &memcache.Item{
		Key: key,
		Value: []byte(value),
		Expiration: expiration,
	}
	return m.mc.Replace(item)
}

// Delete deletes the item with the provided key. The error ErrCacheMiss is
// returned if the item didn't already exist in the cache.
func (m *Memcache) Delete(key string) error {
	return m.mc.Delete(key)
}

// Touch updates the expiry for the given key. The seconds parameter is either
// a Unix timestamp or, if seconds is less than 1 month, the number of seconds
// into the future at which time the item will expire. Zero means the item has
// no expiration time. ErrCacheMiss is returned if the key is not in the cache.
// The key must be at most 250 bytes in length.
func (m *Memcache) Touch(key string, seconds int32) (err error) {
	return m.mc.Touch(key, seconds)
}