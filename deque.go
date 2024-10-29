package gollection

import "slices"

// Deque is an interface for a double-ended queue data structure.
type Deque[T any] interface {
	// Len returns the number of elements in the queue.
	Len() int

	// PopFront removes and returns the element at the front of the queue.
	PopFront() T

	// PopBack removes and returns the element at the back of the queue.
	PopBack() T

	// PushFront adds element to the front of the queue.
	PushFront(...T)

	// PushBack adds element to the back of the queue.
	PushBack(...T)

	// PeekFront returns the element at the front of the queue without removing it.
	PeekFront() T

	// PeekBack returns the element at the back of the queue without removing it.
	PeekBack() T

	// ToSlice returns a copy of the queue as a slice.
	ToSlice() []T

	// Clear removes all elements from the queue.
	Clear()
}

type deque[T any] struct {
	data []T
}

// NewDeque returns a new queue.
func NewDeque[T any](elems ...T) Deque[T] {
	q := &deque[T]{
		data: []T{},
	}

	for _, e := range elems {
		q.PushBack(e)
	}

	return q
}

func (dq *deque[T]) Len() int {
	return len(dq.data)
}

func (dq *deque[T]) PopFront() T {
	if len(dq.data) == 0 {
		return *new(T)
	}

	elem := dq.data[0]
	dq.data = dq.data[1:]
	return elem
}

func (dq *deque[T]) PopBack() T {
	if len(dq.data) == 0 {
		return *new(T)
	}

	elem := dq.data[len(dq.data)-1]
	dq.data = dq.data[:len(dq.data)-1]
	return elem
}

func (dq *deque[T]) PushFront(a ...T) {
	slices.Reverse(a)
	dq.data = append(a, dq.data...)
}

func (dq *deque[T]) PushBack(a ...T) {
	dq.data = append(dq.data, a...)
}

func (dq *deque[T]) PeekFront() T {
	if len(dq.data) == 0 {
		return *new(T)
	}

	return dq.data[0]
}

func (dq *deque[T]) PeekBack() T {
	if len(dq.data) == 0 {
		return *new(T)
	}

	return dq.data[len(dq.data)-1]
}

func (dq *deque[T]) ToSlice() []T {
	cp := make([]T, len(dq.data))
	copy(cp, dq.data)
	return cp
}

func (dq *deque[T]) Clear() {
	dq.data = []T{}
}
