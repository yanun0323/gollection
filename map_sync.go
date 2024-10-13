package gollection

import "sync"

// SyncMap is an interface for a thread-safe map data structure.
type SyncMap[K comparable, V any] interface {
	// Load returns the value stored in the map for a key, or nil if no value is present.
	Clear()

	// Load returns the value stored in the map for a key, or nil if no value is present.
	Clone() map[K]V

	// Delete deletes the value for a key.
	Iter(fn MapIter[K, V])

	// Len returns the number of items in the map.
	Len() int

	// Load returns the value stored in the map for a key, or nil if no value is present.
	Load(key K) (V, bool)

	// LoadAndSet loads the value stored in the map for a key, and sets it to the result of the given function.
	LoadAndSet(key K, fn func(value V) V)

	// Store sets the value for a key.
	Store(key K, value V)

	// Stores stores multiple values.
	Stores(fn func(store func(key K, value V)))

	// Delete deletes the value for a key.
	Delete(key K)
}

type syncMap[K comparable, V any] struct {
	lock *sync.RWMutex
	data map[K]V
}

// NewSyncMap returns a new thread-safe map.
func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return &syncMap[K, V]{
		lock: &sync.RWMutex{},
		data: map[K]V{},
	}
}

func (m *syncMap[K, V]) Load(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	d, ok := m.data[key]
	return d, ok
}

func (m *syncMap[K, V]) Store(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = value
}

func (m *syncMap[K, V]) Stores(fn func(store func(key K, value V))) {
	m.lock.Lock()
	defer m.lock.Unlock()

	action := func(key K, value V) {
		m.data[key] = value
	}

	fn(action)
}

func (m *syncMap[K, V]) LoadAndSet(key K, fn func(value V) V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = fn(m.data[key])
}

func (m *syncMap[K, V]) Iter(fn MapIter[K, V]) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for k, v := range m.data {
		if !fn(k, v) {
			return
		}
	}
}

func (m *syncMap[K, V]) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.data)
}

func (m *syncMap[K, V]) Clone() map[K]V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	c := make(map[K]V, len(m.data))
	for k, v := range m.data {
		c[k] = v
	}
	return c
}

func (m *syncMap[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	clear(m.data)
}

func (m *syncMap[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.data, key)
}
