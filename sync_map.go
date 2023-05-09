package gollection

import (
	"sync"
)

type SyncMap[K comparable, V any] interface {

	// CompareAndDelete deletes the entry for key if its value is equal to old.
	// The old value must be of a comparable type.
	//
	// If there is no current value for key in the map, CompareAndDelete
	// returns false (even if the old value is the nil interface value).
	CompareAndDelete(key K, old V) (deleted bool)

	// CompareAndSwap swaps the old and new values for key
	// if the value stored in the map is equal to old.
	// The old value must be of a comparable type.
	CompareAndSwap(key K, old V, new V) bool

	// Delete deletes the value for a key.
	Delete(key K)

	// Load returns the value stored in the map for a key, or nil/zero value if no
	// value is present.
	// The ok result indicates whether value was found in the map.
	Load(key K) (value V, ok bool)

	// LoadAndDelete deletes the value for a key, returning the previous value if any.
	// The loaded result reports whether the key was present.
	LoadAndDelete(key K) (value V, loaded bool)

	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.
	LoadOrStore(key K, value V) (actual V, loaded bool)

	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	//
	// Range does not necessarily correspond to any consistent snapshot of the Map's
	// contents: no key will be visited more than once, but if the value for any key
	// is stored or deleted concurrently (including by f), Range may reflect any
	// mapping for that key from any point during the Range call. Range does not
	// block other methods on the receiver; even f itself may call any method on m.
	//
	// Range may be O(N) with the number of elements in the map even if f returns
	// false after a constant number of calls.
	Range(func(key K, value V) bool)

	// Iter return a copy of the map.
	Iter() map[K]V

	// Store sets the value for a key.
	Store(key K, value V)

	// Swap swaps the value for a key and returns the previous value if any.
	// The loaded result reports whether the key was present.
	Swap(key K, value V) (previous V, loaded bool)
}

type syncMap[K comparable, V any] struct {
	zero V
	m    sync.Map
}

func NewMap[K comparable, V any]() SyncMap[K, V] {
	return &syncMap[K, V]{
		m: sync.Map{},
	}
}

func (m *syncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}

func (m *syncMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

func (m *syncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *syncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if ok {
		return v.(V), true
	}
	return m.zero, false
}

func (m *syncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, ok := m.m.LoadAndDelete(key)
	if ok {
		return v.(V), true
	}
	return m.zero, false
}

func (m *syncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.m.LoadOrStore(key, value)
	return v.(V), loaded
}

func (m *syncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value interface{}) bool {
		return f(key.(K), value.(V))
	})
}

func (m *syncMap[K, V]) Iter() map[K]V {
	mm := make(map[K]V, 0)
	m.m.Range(func(key, value interface{}) bool {
		mm[key.(K)] = value.(V)
		return true
	})
	return mm
}

func (m *syncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *syncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := m.m.Swap(key, value)
	if loaded {
		return v.(V), true
	}
	return m.zero, false
}
