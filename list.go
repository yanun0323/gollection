package gollection

//Represents a strongly typed list of objects that can be accessed by index. Provides methods to search, sort, and manipulate lists.
//Max stored element quantity = 2,147,483,647
type List struct {
	//Gets the number of elements contained in the List.
	//O(1)
	Count int
	hash  map[int]*interface{}
}

//Initializes a new instance of the List class that is empty and has the default initial capacity.
//O(1)
func NewList() IList {
	return &List{Count: 0, hash: map[int]*interface{}{}}
}

//Adds an object to the end of the List.
//O(1)
func (l *List) ADD(objects ...interface{}) {
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		l.hash[l.Count] = &obj
		l.Count++
	}
}

//Gets the element at the specified index.
//Return nil when the index of List is empty.
//O(1)
func (l *List) At(index int) interface{} {
	if index < 0 || l.IsEmpty() {
		return nil
	}

	obj, exist := l.hash[index]
	if !exist {
		return nil
	}

	return *obj
}

//Removes all elements from the List.
//O(1)
func (l *List) Clear() {
	l.Count = 0
	l.hash = map[int]*interface{}{}
}

////Clone the List without clone the objects inside the List.
//O(n)
func (l *List) Clone() IList {
	m := map[int]*interface{}{}
	for k, v := range l.hash {
		m[k] = v
	}
	return &List{Count: l.Count, hash: m}
}

//Determines whether an element is in the List.
//O(n)
func (l *List) Contains(object interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	for i := 0; i < l.Count; i++ {
		if *l.hash[i] == object {
			return true
		}
	}

	return false
}

//Determines whether any element is in the List.
//O(n)
func (l *List) ContainsAny(objects ...interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	for i := 0; i < l.Count; i++ {

		for _, obj := range objects {
			if *l.hash[i] == obj {
				return true
			}
		}
	}

	return false
}

//Inserts any element into the List at the specified index.
//O(n)
func (l *List) Insert(index int, objects ...interface{}) bool {
	if index < 0 || index > l.Count {
		return false
	}
	if index == l.Count {
		l.ADD(objects...)
		return true
	}

	quantity := len(objects)
	for i := index; i < l.Count; i++ {
		l.hash[i+quantity] = l.hash[i]
	}

	for i := 0; i < quantity; i++ {
		e := objects[i]
		l.hash[i+index] = &e
	}

	l.Count += quantity
	return true
}

//Return true when the List is empty.
//O(1)
func (l *List) IsEmpty() bool {
	return l.Count == 0
}

//Removes the first occurrence of a specific object from the List.
//O(n)
func (l *List) Remove(object interface{}) bool {
	if l.IsEmpty() {
		return true
	}
	found := false
	for i := 0; i < l.Count; i++ {
		if found {
			l.hash[i-1] = l.hash[i]
			continue
		}
		if *l.hash[i] == object {
			found = true
			continue
		}
	}
	l.Count--
	delete(l.hash, l.Count)
	return found
}

//Removes all the elements that match the conditions defined by the specified predicate.
//O(n)
func (l *List) RemoveAll(object interface{}) bool {
	if l.IsEmpty() {
		return true
	}
	found := 0
	for i := 0; i < l.Count; i++ {
		if *l.hash[i] == object {
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

//Removes the element at the specified index of the List.
//O(n)
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

//Sets the element at the specified index.
//Return nil when the index of List is empty.
//O(1)
func (l *List) Set(index int, object interface{}) bool {
	if index < 0 || index > l.Count || l.IsEmpty() {
		return false
	}

	l.hash[index] = &object
	return true
}

//Copies the elements of the List to a new array.
//O(n)
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
