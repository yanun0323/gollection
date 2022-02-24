package datastruct

//max stored element quantity = 2,147,483,647
type Linklist struct {
	Count int
	first *node
	last  *node
}

//BigO(1)
func NewLinklist() Linklist {
	return Linklist{Count: 0, first: nil, last: nil}
}

//BigO(1)
func (l *Linklist) Add(e interface{}) {
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
	return l.first == nil
}
