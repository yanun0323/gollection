package gollection

import "sync"

type syncSet[T comparable] struct {
	rwLock *sync.RWMutex
	set    Set[T]
}

// NewSyncSet returns a new thread-safe set.
func NewSyncSet[T comparable](elems ...T) Set[T] {
	return &syncSet[T]{
		rwLock: &sync.RWMutex{},
		set:    NewSet[T](elems...),
	}
}

func (s *syncSet[T]) Contain(a T) bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.set.Contain(a)
}

func (s *syncSet[T]) Len() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.set.Len()
}
func (s *syncSet[T]) Insert(as ...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.set.Insert(as...)
}

func (s *syncSet[T]) Remove(as ...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.set.Remove(as...)
}

func (s *syncSet[T]) ToSlice() []T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.set.ToSlice()
}

func (s *syncSet[T]) Clear() {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.set.Clear()
}
