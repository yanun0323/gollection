package gollection

import "sync"

type SyncQueue[T any] Queue[T]

func NewSyncQueue[T any]() SyncQueue[T] {
	return &syncQueue[T]{
		rwLock: &sync.RWMutex{},
		q:      NewQueue[T](),
	}
}

type syncQueue[T any] struct {
	rwLock *sync.RWMutex
	q      Queue[T]
}

func (s *syncQueue[T]) Len() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.q.Len()
}

func (s *syncQueue[T]) Dequeue() T {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.q.Dequeue()
}

func (s *syncQueue[T]) Enqueue(...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.q.Enqueue()
}

func (s *syncQueue[T]) Peek() T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.q.Peek()
}

func (s *syncQueue[T]) ToSlice() []T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.q.ToSlice()
}
