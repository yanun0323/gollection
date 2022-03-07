package gollection

//Represents a variable size last-in-first-out (LIFO) collection of instances of the same specified type.
//Max stored object quantity = 2,147,483,647.
type Stack struct {
	count int
	last  *node
}

//Initializes a new instance of the Stack class that is empty and has the default initial capacity.
//O(1)
func NewStack() Stack {
	return Stack{count: 0, last: nil}
}

//Removes all objects from the Stack.
//O(1)
func (s *Stack) Clear() {
	s.count = 0
	s.last = nil
}

//Clone the Stack without clone the objects inside the Stack.
//O(1)
func (s *Stack) Clone() Stack {
	return Stack{count: s.count, last: s.last}
}

//Determines whether any element is in the Stack.
//O(n)
func (s *Stack) Contains(objects ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}

	node := s.last
	for i := 0; i < s.count; i++ {
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
	return false
}

//Gets the number of elements contained in the Stack.
//O(1)
func (s *Stack) Count() int {
	return s.count
}

//Return true when the Stack is empty.
//O(1)
func (s *Stack) IsEmpty() bool {
	return s.last == nil
}

//Returns the object at the beginning of the Stack without removing it.
//O(1)
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return *s.last.data
}

//Removes and returns the object at the top of the Stack.
//Return nil when the Stack is empty.
//O(1)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	node := s.last
	s.last = node.forward
	if s.count > 0 {
		s.count--
	}
	return *node.data
}

//Inserts an object at the top of the Stack.
//O(1)
func (s *Stack) Push(object interface{}) {
	node := newNode(&object, nil, s.last)
	s.last = node
	s.count++
}

//Copies the Stack to a new slice.
//O(n)
func (s *Stack) ToArray() []interface{} {
	if s.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, s.count)

	node := s.last
	for i := 0; i < s.count; i++ {
		arr[i] = *node.data
		node = node.forward
	}
	return arr
}
