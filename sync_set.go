package gollection

import "sync"

type SyncSet[T comparable] interface {
	Contain(T) bool
	Len() int
	Insert(...T)
	Remove(...T)
	// Iter return a copy of the set.
	Iter() []T
}

type syncSet[T comparable] struct {
	mu sync.Mutex
	m  map[T]struct{}
}

func NewSyncSet[T comparable]() SyncSet[T] {
	return &syncSet[T]{
		m: map[T]struct{}{},
	}
}

func (s *syncSet[T]) Contain(a T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.m[a]
	return ok
}

func (s *syncSet[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for range s.m {
		count++
	}
	return count
}
func (s *syncSet[T]) Insert(as ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, a := range as {
		s.m[a] = struct{}{}
	}
}

func (s *syncSet[T]) Remove(as ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, a := range as {
		delete(s.m, a)
	}
}

func (s *syncSet[T]) Iter() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	mm := []T{}
	for k := range s.m {
		mm = append(mm, k)
	}
	return mm
}
