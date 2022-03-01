package gollection

//max stored element quantity = 2,147,483,647
type Linkedlist struct {
	Count int
	first *node
	last  *node
}

//BigO(1)
func NewLinkedlist() Linkedlist {
	return Linkedlist{Count: 0, first: nil, last: nil}
}

//Return first data without remove it, BigO(1)
func (l *Linkedlist) First() interface{} {
	if l.IsEmpty() {
		return nil
	}

	return *l.first.data
}

//Return last data without remove it, BigO(1)
func (l *Linkedlist) Last() interface{} {
	if l.IsEmpty() {
		return nil
	}

	if l.last == nil {
		return *l.first.data
	}

	return *l.last.data
}

// // BigO(1)
// func (l *Linkedlist) AddFirst(e interface{}) {
// 	n := newNode(&e, nil, nil)
// 	if l.IsEmpty() {
// 		l.first = n
// 		l.Count++
// 		return
// 	}

// 	if l.last == nil {
// 		l.last = l.first
// 	}

// 	l.first.forward = n
// 	n.backward = l.first
// 	l.first = n
// 	l.Count++
// }

// //BigO(1)
// func (l *Linkedlist) AddLast(e interface{}) {
// 	n := newNode(&e, nil, nil)

// 	if l.IsEmpty() {
// 		l.AddFirst(e)
// 		return
// 	}

// 	l.last.backward = n
// 	n.forward = l.last
// 	l.last = n
// 	l.Count++
// }

// //Return first data and remove it, BigO(1)
// func (l *Linkedlist) TakeFirst() interface{} {
// 	if l.IsEmpty() {
// 		return nil
// 	}

// 	n := l.first
// 	l.first = n.backward
// 	l.Count--

// 	if l.first.backward == l.first {
// 		l.first.backward = nil
// 		l.last = nil
// 	}

// 	return *n.data
// }

// //Return last data and remove it, BigO(1)
// func (l *Linkedlist) TakeLast() interface{} {
// 	if l.IsEmpty() {
// 		return nil
// 	}

// 	if l.last == nil {
// 		n := l.first
// 		l.first = n.backward
// 		l.Count--
// 		return *n.data
// 	}

// 	n := l.last
// 	l.last = n.forward
// 	l.Count--

// 	return *n.data
// }

//BigO(1)
func (l *Linkedlist) IsEmpty() bool {
	return l.first == nil && l.last == nil
}

//BigO(1)
func (l *Linkedlist) Clear() {
	l.Count = 0
	l.last = nil
	l.first = nil
}

//BigO(1)
func (l *Linkedlist) Clone() Linkedlist {
	return Linkedlist{
		Count: l.Count,
		first: l.first,
		last:  l.last}
}

//Return true if there's any element in queue, BigO(n)
func (l *Linkedlist) Contains(es ...interface{}) bool {
	if l.IsEmpty() {
		return false
	}

	n := l.first
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

//BigO(n)
func (l *Linkedlist) ToArray() []interface{} {
	if l.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, l.Count)

	n := l.first
	for i := 0; i < l.Count; i++ {
		arr[i] = *n.data
		n = n.backward
	}
	return arr
}
