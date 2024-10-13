package gollection

// Queue is an interface for a queue data structure.
type Queue[T any] interface {
	// Len returns the number of elements in the queue.
	Len() int

	// Dequeue removes and returns the element at the front of the queue.
	Dequeue() T

	// Enqueue adds element to the back of the queue.
	Enqueue(...T)

	// Peek returns the element at the front of the queue without removing it.
	Peek() T

	// ToSlice returns a copy of the queue as a slice.
	ToSlice() []T
}

type queue[T any] struct {
	zero T
	data []T
}

// NewQueue returns a new queue.
func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		data: []T{},
	}
}

func (q *queue[T]) Enqueue(a ...T) {
	q.data = append(q.data, a...)
}

func (q *queue[T]) Dequeue() T {
	if len(q.data) == 0 {
		return q.zero
	}
	elem := q.data[0]
	q.data = q.data[1:]
	return elem
}

func (q *queue[T]) Len() int {
	return len(q.data)
}

func (q *queue[T]) Peek() T {
	if len(q.data) == 0 {
		return q.zero
	}
	return q.data[0]
}

func (q *queue[T]) ToSlice() []T {
	return append(make([]T, 0, len(q.data)), q.data...)
}
