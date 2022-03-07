package gollection

//Represents a first-in, first-out collection of objects.
//max stored element quantity = 2,147,483,647
type Queue struct {
	count int
	first *node
	last  *node
}

//Initializes a new instance of the Queue class that is empty and has the default initial capacity.
//O(1)
func NewQueue() IQueue {
	return &Queue{count: 0, first: nil, last: nil}
}

//Removes all objects from the Queue.
//O(1)
func (q *Queue) Clear() bool {
	q.count = 0
	q.first = nil
	q.last = nil
	return true
}

//Clone the Queue without clone the objects inside the Queue.
//O(1)
func (q *Queue) Clone() IQueue {
	return &Queue{
		count: q.count,
		first: q.first,
		last:  q.last}
}

//Determines whether any element is in the Queue.
//O(n)
func (q *Queue) Contains(objects ...interface{}) bool {
	if q.IsEmpty() {
		return false
	}

	node := q.first
	for i := 0; i < q.count; i++ {
		for _, obj := range objects {
			if *node.data == obj {
				return true
			}
		}
		node = node.backward
		if node == nil {
			return false
		}
	}
	return false
}

//Gets the number of elements contained in the Queue.
//O(1)
func (q *Queue) Count() int {
	return q.count
}

//Removes and returns the object at the beginning of the Queue.
//Return false when the Queue is empty.
//O(1)
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	node := q.first
	q.first = node.backward
	if q.count > 0 {
		q.count--
	}
	return *node.data, true
}

//Adds an object to the end of the Queue.
//O(1)
func (q *Queue) Enqueue(object interface{}) bool {
	node := newNode(&object, nil, nil)

	if q.IsEmpty() {
		q.first = node
		q.first.backward = node
		q.last = node
		q.count++
		return true
	}

	q.last.backward = node
	q.last = node
	q.count++
	return true
}

//Return true when the Queue is empty.
//O(1)
func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

//Returns the object at the beginning of the Queue without removing it.
//O(1)
func (q *Queue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return *q.first.data, true
}

//Copies the Queue to a new slice.
//Return false when the Queue is empty.
//O(n)
func (q *Queue) ToArray() ([]interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	arr := make([]interface{}, q.count)

	node := q.first
	for i := 0; i < q.count; i++ {
		arr[i] = *node.data
		node = node.backward
	}
	return arr, true
}
