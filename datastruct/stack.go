package datastruct

//max stored element quantity = 2,147,483,647
type Stack struct {
	Count int
	last  *nodeF
}

//BigO(1)
func NewStack() Stack {
	return Stack{Count: 0, last: nil}
}

//BigO(1)
func (s *Stack) Push(e interface{}) {
	n := newNodeF(&e, s.last)
	s.last = n
	s.Count++
}

//BigO(1)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	result := s.last
	s.last = result.forward
	if s.Count > 0 {
		s.Count--
	}
	return *result.data
}

//BigO(1)
func (s *Stack) IsEmpty() bool {
	return s.last == nil
}

//BigO(1)
func (s *Stack) Clear() {
	s.Count = 0
	s.last = nil
}

//clone a stack but not the elements, BigO(1)
func (s *Stack) Clone() Stack {
	return Stack{Count: s.Count, last: s.last}
}

//Return true if there's one element in stack, BigO(n)
func (s *Stack) Contains(es ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}

	n := s.last
	for {
		for _, e := range es {
			if *n.data == e {
				return true
			}
		}
		n = n.forward
		if n == nil {
			return false
		}
	}
}

//BigO(1)
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return *s.last.data
}

//BigO(n)
func (s *Stack) ToArray() []interface{} {
	if s.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, s.Count)

	n := s.last
	for i := 0; i < s.Count; i++ {
		arr[i] = *n.data
		n = n.forward
	}
	return arr
}
