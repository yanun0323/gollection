package gollection

import "sync"

type SyncStack[T any] Stack[T]

func NewSyncStack[T any]() SyncStack[T] {
	return &syncStack[T]{
		rwLock: &sync.RWMutex{},
		stack:  NewStack[T](),
	}
}

type syncStack[T any] struct {
	rwLock *sync.RWMutex
	stack  Stack[T]
}

func (s *syncStack[T]) Len() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.stack.Len()
}

func (s *syncStack[T]) Peek() T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.stack.Peek()
}

func (s *syncStack[T]) Pop() T {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.stack.Pop()
}

func (s *syncStack[T]) Push(as ...T) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.stack.Push(as...)
}

func (s *syncStack[T]) ToSlice() []T {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.stack.ToSlice()
}
