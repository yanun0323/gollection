package gollection

//max stored element quantity = 2,147,483,647
type List struct {
	Count int
	hash  map[int]*interface{}
}

//BigO(1)
func NewList() List {
	return List{Count: 0, hash: map[int]*interface{}{}}
}

//Return nil if find nothing, BigO(1)
func (l *List) At(index int) interface{} {
	if index < 0 || l.IsEmpty() {
		return nil
	}

	e, exist := l.hash[index]
	if !exist {
		return nil
	}

	return *e
}

//BigO(1)
func (l *List) ADD(es ...interface{}) {
	for i := 0; i < len(es); i++ {
		e := es[i]
		l.hash[l.Count] = &e
		l.Count++
	}
}

//BigO(1)
func (l *List) Clear() {
	l.Count = 0
	l.hash = map[int]*interface{}{}
}

//BigO(n)
func (l *List) Contains(e interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	for i := 0; i < l.Count; i++ {
		if *l.hash[i] == e {
			return true
		}
	}

	return false
}

//BigO(n)
func (l *List) Insert(index int, es ...interface{}) bool {
	if index < 0 || index > l.Count {
		return false
	}
	if index == l.Count {
		l.ADD(es...)
		return true
	}

	quantity := len(es)
	for i := index; i < l.Count; i++ {
		l.hash[i+quantity] = l.hash[i]
	}

	for i := 0; i < quantity; i++ {
		e := es[i]
		l.hash[i+index] = &e
	}

	l.Count += quantity
	return true
}

//BigO(n)
func (l *List) ToArray() []interface{} {
	if l.IsEmpty() {
		return nil
	}

	arr := make([]interface{}, l.Count)
	for i := 0; i < l.Count; i++ {
		arr[i] = *l.hash[i]
	}
	return arr
}

//BigO(n)
func (l *List) RemoveAt(index int) bool {
	if index < 0 || index >= l.Count {
		return false
	}
	for i := index + 1; i < l.Count; i++ {
		l.hash[i-1] = l.hash[i]
	}
	l.Count--
	delete(l.hash, l.Count)
	return true
}

//BigO(n)
func (l *List) Remove(e interface{}) bool {
	if l.IsEmpty() {
		return true
	}
	found := false
	for i := 0; i < l.Count; i++ {
		if found {
			l.hash[i-1] = l.hash[i]
			continue
		}
		if *l.hash[i] == e {
			found = true
			continue
		}
	}
	l.Count--
	delete(l.hash, l.Count)
	return found
}

//BigO(n)
func (l *List) RemoveAll(e interface{}) bool {
	if l.IsEmpty() {
		return true
	}
	found := 0
	for i := 0; i < l.Count; i++ {
		if *l.hash[i] == e {
			found++
			continue
		}
		if found != 0 {
			l.hash[i-found] = l.hash[i]
			continue
		}
	}
	l.Count -= found
	for i := 0; i < found; i++ {
		delete(l.hash, l.Count+i)
	}
	return found != 0
}

//BigO(1)
func (l *List) IsEmpty() bool {
	return l.Count == 0
}

//Clone a list but not the elements, BigO(n)
func (l *List) Clone() List {
	m := map[int]*interface{}{}
	for k, v := range l.hash {
		m[k] = v
	}
	return List{Count: l.Count, hash: m}
}
