package gollection

type Stack[T any] interface {
	Count() int
	Peek() T
	Pop() T
	Push(...T)
	ToSlice() []T
}

type stack[T any] struct {
	zero T
	data []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(a ...T) {
	s.data = append(s.data, a...)
}

func (s *stack[T]) Pop() T {
	if len(s.data) == 0 {
		return s.zero
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem
}

func (s *stack[T]) Peek() T {
	if len(s.data) == 0 {
		return s.zero
	}
	return s.data[len(s.data)-1]
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack[T]) ToSlice() []T {
	result := make([]T, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}
