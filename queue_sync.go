package gollection

import "sync"

type syncQueue[T any] struct {
	rwLock *sync.RWMutex
	q      queue[T]
}

// NewSyncQueue returns a new thread-safe queue.
func NewSyncQueue[T any](elems ...T) Queue[T] {
	return &syncQueue[T]{
		rwLock: &sync.RWMutex{},
		q:      queue[T]{data: elems},
	}
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

func (s *syncQueue[T]) Enqueue(elems ...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.q.Enqueue(elems...)
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

func (s *syncQueue[T]) Clear() {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.q.Clear()
}

func (s *syncQueue[T]) Shrink(n uint) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.q.Shrink(n)
}
