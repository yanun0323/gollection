package gollection

import "sync"

func NewSyncSet[T comparable]() Set[T] {
	return &syncSet[T]{
		rwLock: &sync.RWMutex{},
		set:    NewSet[T](),
	}
}

type syncSet[T comparable] struct {
	rwLock *sync.RWMutex
	set    Set[T]
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
