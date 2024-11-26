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

	// Clear removes all elements from the queue.
	Clear()

	// Shrink shrinks the length of the queue to the given length.
	Shrink(uint)
}

type queue[T any] struct {
	zero T
	data []T
}

// NewQueue returns a new queue.
func NewQueue[T any](elems ...T) Queue[T] {
	q := &queue[T]{
		data: []T{},
	}

	for _, e := range elems {
		q.Enqueue(e)
	}

	return q
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

func (q *queue[T]) Clear() {
	q.data = []T{}
}

func (q *queue[T]) Shrink(n uint) {
	if n == 0 {
		q.Clear()
		return
	}

	if len(q.data) <= int(n) {
		return
	}

	q.data = q.data[:n]
}
