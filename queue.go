package gollection

type Queue[T any] interface {
	Count() int
	Dequeue() T
	Enqueue(...T)
	Peek() T
	ToSlice() []T
}

type queue[T any] struct {
	zero T
	data []T
}

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

func (q *queue) Len() int {
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
