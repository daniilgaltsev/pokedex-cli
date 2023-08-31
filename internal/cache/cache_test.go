package cache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(time.Duration(5 * float64(time.Millisecond)))
	to_add := []struct{
		key string
		value []byte
	}{
		{"test", []byte("test")},
		{"test2", []byte("test2")},
	}

	cache.Add(to_add[0].key, to_add[0].value)
	value, ok := cache.Get(to_add[0].key)
	if !ok {
		t.Errorf("cache.Get(%v) == %v, expected %v", to_add[0].key, ok, true)
		return
	}
	if string(value) != string(to_add[0].value) {
		t.Errorf("cache.Get(%v) == %v, expected %v", to_add[0].key, string(value), string(to_add[0].value))
		return
	}
	_, ok = cache.Get(to_add[1].key)
	if ok {
		t.Errorf("cache.Get(%v) == %v, expected %v", to_add[1].key, ok, false)
		return
	}

	cache.Add("test2", []byte("test2"))
	
	for _, next := range to_add {
		value, ok := cache.Get(next.key)
		if !ok {
			t.Errorf("cache.Get(%v) == %v, expected %v", next.key, ok, true)
			return
		}
		if string(value) != string(next.value) {
			t.Errorf("cache.Get(%v) == %v, expected %v", next.key, string(value), string(next.value))
			return
		}
	}
}

func TestCacheRetention(t *testing.T) {
	retention := time.Duration(5 * float64(time.Millisecond))
	wait := retention + time.Duration(5 * float64(time.Millisecond))

	cache := NewCache(retention)
	cache.Add("test", []byte("test"))

	_, ok := cache.Get("test")
	if !ok {
		t.Errorf("cache.Get(%v) == %v, expected %v", "test", ok, true)
		return
	}

	time.Sleep(wait) // NOTE: don't really like waiting in tests

	_, ok = cache.Get("test")
	if ok {
		t.Errorf("cache.Get(%v) == %v, expected %v", "test", ok, false)
		return
	}
	
}
