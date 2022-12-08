package gollection

type stack struct {
	data []any
}

func NewStack() stack {
	return stack{
		data: []any{},
	}
}

func (s *stack) Push(a ...any) {
	s.data = append(s.data, a...)
}

func (s *stack) Pop() any {
	if len(s.data) == 0 {
		return nil
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem
}

func (s *stack) Peek() any {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) ToSlice() []any {
	result := make([]any, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}
