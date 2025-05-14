package gollection

import "sync"

// Slice is an interface for a slice data structure.
type Slice[T any] interface {
	// Len returns the number of elements in the slice.
	Len() int

	// Get returns the element at the index.
	Get(int) (T, bool)

	// First returns the first element of the slice.
	First() (T, bool)

	// Last returns the last element of the slice.
	Last() (T, bool)

	// Set sets the element at the index.
	Set(int, T) bool

	// Swap swaps the elements at the indices.
	Swap(int, T) bool

	// Append appends elements to the slice.
	Append(...T)

	// ToSlice returns a copy of the slice as a slice.
	ToSlice() []T

	// Shrink shrinks the length of the slice to the given length.
	Shrink(uint)

	// Iter iterates over the slice.
	Iter(Iter[T])

	// Update updates the slice.
	Update(func([]T) []T)
}

type syncSlice[T any] struct {
	lock *sync.RWMutex
	data []T
}

func NewSyncSlice[T any](elems ...T) Slice[T] {
	slice := []T{}
	if len(elems) != 0 {
		slice = elems
	}

	return &syncSlice[T]{
		lock: &sync.RWMutex{},
		data: slice,
	}
}

func (s *syncSlice[T]) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.data)
}

func (s *syncSlice[T]) Get(index int) (T, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if index >= len(s.data) {
		return *new(T), false
	}

	return s.data[index], true
}

func (s *syncSlice[T]) First() (T, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if len(s.data) == 0 {
		return *new(T), false
	}

	return s.data[0], true
}

func (s *syncSlice[T]) Last() (T, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if len(s.data) == 0 {
		return *new(T), false
	}

	return s.data[len(s.data)-1], true
}

func (s *syncSlice[T]) Set(index int, elem T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if index >= len(s.data) {
		return false
	}

	s.data[index] = elem

	return true
}

func (s *syncSlice[T]) Swap(index int, elem T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if index >= len(s.data) {
		return false
	}

	s.data[index] = elem

	return true
}

func (s *syncSlice[T]) Append(elems ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, elems...)
}

func (s *syncSlice[T]) ToSlice() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	copied := make([]T, len(s.data))
	copy(copied, s.data)

	return copied
}

func (s *syncSlice[T]) Shrink(length uint) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if length >= uint(len(s.data)) {
		return
	}

	s.data = s.data[:length]
}

func (s *syncSlice[T]) Iter(iter Iter[T]) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, elem := range s.data {
		if iter(elem) {
			continue
		}

		return
	}
}

func (s *syncSlice[T]) Update(update func([]T) []T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = update(s.data)
}
