package gollection

import "sync"

type syncPriorityQueue[T any] struct {
	rwLock *sync.RWMutex
	pq     Queue[T]
}

// NewSyncPriorityQueue returns a new thread-safe priority queue.
func NewSyncPriorityQueue[T any](score func(T, T) bool, elems ...T) Queue[T] {
	return &syncPriorityQueue[T]{
		rwLock: &sync.RWMutex{},
		pq:     NewPriorityQueue[T](score, elems...),
	}
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

func (s *syncPriorityQueue[T]) Clear() {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.pq.Clear()
}
