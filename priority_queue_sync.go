package gollection

import "sync"

func NewSyncPriorityQueue[T any](score func(T, T) bool) Queue[T] {
	return &syncPriorityQueue[T]{
		rwLock: &sync.RWMutex{},
		pq:     NewPriorityQueue[T](score),
	}
}

type syncPriorityQueue[T any] struct {
	rwLock *sync.RWMutex
	pq     Queue[T]
}

func (s *syncPriorityQueue[T]) Len() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.pq.Len()
}

func (s *syncPriorityQueue[T]) Dequeue() T {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.pq.Dequeue()
}

func (s *syncPriorityQueue[T]) Enqueue(as ...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.pq.Enqueue(as...)
}

func (s *syncPriorityQueue[T]) Peek() T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.pq.Peek()
}

func (s *syncPriorityQueue[T]) ToSlice() []T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.pq.ToSlice()
}
