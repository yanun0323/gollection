package gollection

import "sync"

type syncDeque[T any] struct {
	mu *sync.RWMutex
	d  Deque[T]
}

func NewSyncDeque[T any](elems ...T) Deque[T] {
	return &syncDeque[T]{
		mu: &sync.RWMutex{},
		d:  NewDeque(elems...),
	}
}

func (dq *syncDeque[T]) Clear() {
	dq.mu.Lock()
	defer dq.mu.Unlock()

	dq.d.Clear()
}

func (dq *syncDeque[T]) Len() int {
	dq.mu.RLock()
	defer dq.mu.RUnlock()

	return dq.d.Len()
}

func (dq *syncDeque[T]) PeekBack() T {
	dq.mu.RLock()
	defer dq.mu.RUnlock()

	return dq.d.PeekBack()
}

func (dq *syncDeque[T]) PeekFront() T {
	dq.mu.RLock()
	defer dq.mu.RUnlock()

	return dq.d.PeekFront()
}

func (dq *syncDeque[T]) PopBack() T {
	dq.mu.Lock()
	defer dq.mu.Unlock()

	return dq.d.PopBack()
}

func (dq *syncDeque[T]) PopFront() T {
	dq.mu.Lock()
	defer dq.mu.Unlock()

	return dq.d.PopFront()
}

func (dq *syncDeque[T]) PushBack(a ...T) {
	dq.mu.Lock()
	defer dq.mu.Unlock()

	dq.d.PushBack(a...)
}

func (dq *syncDeque[T]) PushFront(a ...T) {
	dq.mu.Lock()
	defer dq.mu.Unlock()

	dq.d.PushFront(a...)
}

func (dq *syncDeque[T]) ToSlice() []T {
	dq.mu.RLock()
	defer dq.mu.RUnlock()

	return dq.d.ToSlice()
}
