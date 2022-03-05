package gollection

//Represents a variable size last-in-first-out (LIFO) collection of instances of the same specified type.
//Max stored object quantity = 2,147,483,647.
type Stack struct {
	//Gets the number of elements contained in the Stack.
	//O(1)
	Count int
	last  *node
}

//Initializes a new instance of the Stack class that is empty and has the default initial capacity.
//O(1)
func NewStack() IStack {
	return &Stack{Count: 0, last: nil}
}

//Removes all objects from the Stack.
//O(1)
func (s *Stack) Clear() {
	s.Count = 0
	s.last = nil
}

//Clone the Stack without clone the objects inside the Stack.
//O(1)
func (s *Stack) Clone() IStack {
	return &Stack{Count: s.Count, last: s.last}
}

//Determines whether an element is in the Stack.
//O(n)
func (s *Stack) Contains(object T) bool {
	if s.IsEmpty() {
		return false
	}

	node := s.last
	for {
		if *node.data == object {
			return true
		}
		node = node.forward
		if node == nil {
			return false
		}
	}
}

//Determines whether any element is in the Stack.
//O(n)
func (s *Stack) ContainsAny(objects ...T) bool {
	if s.IsEmpty() {
		return false
	}

	node := s.last
	for {
		for _, obj := range objects {
			if *node.data == obj {
				return true
			}
		}
		node = node.forward
		if node == nil {
			return false
		}
	}
}

//Return true when the Stack is empty.
//O(1)
func (s *Stack) IsEmpty() bool {
	return s.last == nil
}

//Returns the object at the beginning of the Stack without removing it.
//O(1)
func (s *Stack) Peek() T {
	if s.IsEmpty() {
		return nil
	}
	return *s.last.data
}

//Removes and returns the object at the top of the Stack.
//Return nil when the Stack is empty.
//O(1)
func (s *Stack) Pop() T {
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

//Inserts an object at the top of the Stack.
//O(1)
func (s *Stack) Push(object T) {
	node := newNode(&object, nil, s.last)
	s.last = node
	s.Count++
}

//Copies the Stack to a new array.
//O(n)
func (s *Stack) ToArray() []T {
	if s.IsEmpty() {
		return nil
	}
	arr := make([]T, s.Count)

	node := s.last
	for i := 0; i < s.Count; i++ {
		arr[i] = *node.data
		node = node.forward
	}
	return arr
}
