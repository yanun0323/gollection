package gollection

//Represents a first-in, first-out collection of objects.
//max stored element quantity = 2,147,483,647
type Queue struct {
	//Gets the number of elements contained in the Queue.
	//O(1)
	Count int
	first *node
	last  *node
}

//Initializes a new instance of the Queue class that is empty and has the default initial capacity.
//O(1)
func NewQueue() IQueue {
	return &Queue{Count: 0, first: nil, last: nil}
}

//Removes all objects from the Queue.
//O(1)
func (q *Queue) Clear() {
	q.Count = 0
	q.last = nil
	q.first = nil
}

//Clone the Queue without clone the objects inside the Queue.
//O(1)
func (q *Queue) Clone() IQueue {
	return &Queue{
		Count: q.Count,
		first: q.first,
		last:  q.last}
}

//Determines whether an element is in the Queue.
//O(n)
func (q *Queue) Contains(object interface{}) bool {
	if q.IsEmpty() {
		return false
	}

	node := q.first
	for {
		if *node.data == object {
			return true
		}
		node = node.backward
		if node == nil {
			return false
		}
	}
}

//Determines whether any element is in the Queue.
//O(n)
func (q *Queue) ContainsAny(objects ...interface{}) bool {
	if q.IsEmpty() {
		return false
	}

	node := q.first
	for {
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
}

//Removes and returns the object at the beginning of the Queue.
//Return nil when the Queue is empty.
//O(1)
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	result := q.first
	q.first = result.backward
	if q.Count > 0 {
		q.Count--
	}
	return *result.data
}

//Adds an object to the end of the Queue.
//O(1)
func (q *Queue) Enqueue(e interface{}) {
	node := newNode(&e, nil, nil)

	if q.IsEmpty() {
		q.first = node
		q.first.backward = node
		q.last = node
		q.Count++
		return
	}

	q.last.backward = node
	q.last = node
	q.Count++
}

//Return true when the Queue is empty.
//O(1)
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

//Returns the object at the beginning of the Queue without removing it.
//O(1)
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return *q.first.data
}

//Copies the Queue to a new array.
//O(n)
func (q *Queue) ToArray() []interface{} {
	if q.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, q.Count)

	node := q.first
	for i := 0; i < q.Count; i++ {
		arr[i] = *node.data
		node = node.backward
	}
	return arr
}
