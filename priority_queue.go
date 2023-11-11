package gollection

import "container/heap"

type PriorityQueue[T any] interface {
	Len() int
	Dequeue() T
	Enqueue(...T)
	Peek() T
	ToSlice() []T
}

func NewPriorityQueue[T any](greater func(T, T) bool) PriorityQueue[T] {
	return &priorityQueue[T]{
		data:    []T{},
		greater: greater,
	}
}

type priorityQueue[T any] struct {
	zero    T
	data    []T
	greater func(T, T) bool
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.data)
}

func (pq *priorityQueue[T]) Dequeue() T {
	if len(pq.data) == 0 {
		return pq.zero
	}
	return heap.Pop(pq).(T)
}

func (pq *priorityQueue[T]) Enqueue(as ...T) {
	for _, a := range as {
		heap.Push(pq, a)
	}
}

func (pq *priorityQueue[T]) Peek() T {
	if len(pq.data) == 0 {
		return pq.zero
	}
	return pq.data[0]
}

func (pq *priorityQueue[T]) ToSlice() []T {
	return append(make([]T, 0, len(pq.data)), pq.data...)
}

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return !pq.greater(pq.data[i], pq.data[j])
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityQueue[T]) Push(x interface{}) {
	pq.data = append(pq.data, x.(T))
}

func (pq *priorityQueue[T]) Pop() interface{} {
	x := pq.data[pq.Len()-1]
	pq.data = pq.data[:pq.Len()-1]
	return x
}
