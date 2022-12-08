package gollection

type queue struct {
	data []any
}

func NewQueue() queue {
	return queue{
		data: []any{},
	}
}

func (q *queue) Enqueue(a ...any) {
	q.data = append(q.data, a...)
}

func (q *queue) Dequeue() any {
	if len(q.data) == 0 {
		return nil
	}
	elem := q.data[0]
	q.data = q.data[1:]
	return elem
}

func (q *queue) Len() int {
	return len(q.data)
}

func (q *queue) Peek() any {
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}

func (q *queue) ToSlice() []any {
	return append(make([]any, 0, len(q.data)), q.data...)
}
