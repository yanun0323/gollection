package gollection

// Stack is an interface for a stack data structure.
type Stack[T any] interface {
	// Len returns the number of elements in the stack.
	Len() int

	// Peek returns the element at the top of the stack without removing it.
	Peek() T

	// Pop removes and returns the element at the top of the stack.
	Pop() T

	// Push adds element to the top of the stack.
	Push(...T)

	// ToSlice returns a copy of the stack as a slice.
	ToSlice() []T
}

type stack[T any] struct {
	zero T
	data []T
}

// NewStack returns a new stack.
func NewStack[T any](elems ...T) Stack[T] {
	s := &stack[T]{
		data: []T{},
	}

	for _, e := range elems {
		s.Push(e)
	}

	return s
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

func (s *stack[T]) Len() int {
	return len(s.data)
}

func (s *stack[T]) ToSlice() []T {
	result := make([]T, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}
