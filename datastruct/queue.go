package datastruct

//max stored element quantity = 2,147,483,647
type Queue struct {
	Count int
	first *nodeB
	last  *nodeB
}

//BigO(1)
func NewQueue() Queue {
	return Queue{Count: 0, first: nil, last: nil}
}

//BigO(1)
func (q *Queue) Enqueue(e interface{}) {
	n := newNodeB(&e, nil)

	if q.IsEmpty() {
		q.first = n
		q.first.backward = n
		q.last = n
		q.Count++
		return
	}

	q.last.backward = n
	q.last = n
	q.Count++
}

//Return nil if queue empty, BigO(1)
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

//BigO(1)
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

//BigO(1)
func (q *Queue) Clear() {
	q.Count = 0
	q.last = nil
	q.first = nil
}

//clone a queue but not the elements, BigO(1)
func (q *Queue) Clone() Queue {
	return Queue{Count: q.Count, first: q.first, last: q.last}
}

//Return true if there's one element in queue, BigO(n)
func (q *Queue) Contains(es ...interface{}) bool {
	if q.IsEmpty() {
		return false
	}

	n := q.first
	for {
		for _, e := range es {
			if *n.data == e {
				return true
			}
		}
		n = n.backward
		if n == nil {
			return false
		}
	}
}

//BigO(1)
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return *q.first.data
}

//BigO(n)
func (q *Queue) ToArray() []interface{} {
	if q.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, q.Count)

	n := q.first
	for i := 0; i < q.Count; i++ {
		arr[i] = *n.data
		n = n.backward
	}
	return arr
}
