package gollection

//max stored element quantity = 2,147,483,647
type Linklist struct {
	Count int
	hash  map[interface{}][]*node
	first *node
	last  *node
}

//BigO(1)
func NewLinklist() Linklist {
	return Linklist{Count: 0, first: nil, last: nil}
}

//Return first data without remove it, BigO(1)
func (l *Linklist) First() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return *l.first.data
}

//Return last data without remove it, BigO(1)
func (l *Linklist) Last() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return *l.last.data
}

//Return first data and remove it, BigO(1)
func (l *Linklist) TakeFirst() interface{} {
	if l.IsEmpty() {
		return nil
	}

	n := l.first
	next := n.backward
	l.first = next
	if next == nil {
		l.last = nil
	}

	return *n.data
}

//Return last data and remove it, BigO(1)
func (l *Linklist) TakeLast() interface{} {
	if l.IsEmpty() {
		return nil
	}

	n := l.last
	next := n.forward
	l.last = next
	if next == nil {
		l.first = nil
	}

	return *n.data
}

//BigO(1)
func (l *Linklist) AddFirst(e interface{}) {
	n := newNode(0, &e, nil, nil)

	if l.IsEmpty() {
		l.last = n
		l.last.forward = n
		l.first = n
		l.Count++
		return
	}

	l.first.backward = n
	l.first = n
	l.Count++
}

//BigO(1)
func (l *Linklist) AddLast(e interface{}) {
	n := newNode(0, &e, nil, nil)

	if l.IsEmpty() {
		l.first = n
		l.first.backward = n
		l.last = n
		l.Count++
		return
	}

	l.last.backward = n
	l.last = n
	l.Count++
}

func (l *Linklist) IsEmpty() bool {
	return l.first == nil && l.last == nil
}
