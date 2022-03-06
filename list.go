package gollection

//Represents a strongly typed list of objects that can be accessed by index. Provides methods to search, sort, and manipulate lists.
//Max stored element quantity = 2,147,483,647
type List struct {
	count int
	hash  map[int]*T
}

//Initializes a new instance of the List class that is empty and has the default initial capacity.
//O(1)
func NewList() IList {
	return &List{count: 0, hash: map[int]*T{}}
}

//Adds an object to the end of the List.
//O(1)
func (l *List) ADD(objects ...T) {
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		l.hash[l.count] = &obj
		l.count++
	}
}

//Gets the element at the specified index.
//Return nil when the index of List is empty.
//O(1)
func (l *List) At(index int) T {
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
	l.count = 0
	l.hash = map[int]*T{}
}

////Clone the List without clone the objects inside the List.
//O(n)
func (l *List) Clone() IList {
	m := map[int]*T{}
	for k, v := range l.hash {
		m[k] = v
	}
	return &List{count: l.count, hash: m}
}

//Determines whether an element is in the List.
//O(n)
func (l *List) Contains(object T) bool {
	if l.IsEmpty() {
		return false
	}
	for i := 0; i < l.count; i++ {
		if *l.hash[i] == object {
			return true
		}
	}

	return false
}

//Determines whether any element is in the List.
//O(n)
func (l *List) ContainsAny(objects ...T) bool {
	if l.IsEmpty() {
		return false
	}
	for i := 0; i < l.count; i++ {

		for _, obj := range objects {
			if *l.hash[i] == obj {
				return true
			}
		}
	}

	return false
}

//Gets the number of elements contained in the List.
//O(1)
func (l *List) Count() int {
	return l.count
}

//Inserts any element into the List at the specified index.
//O(n)
func (l *List) Insert(index int, objects ...T) bool {
	if index < 0 || index > l.count {
		return false
	}
	if index == l.count {
		l.ADD(objects...)
		return true
	}

	quantity := len(objects)
	for i := index; i < l.count; i++ {
		l.hash[i+quantity] = l.hash[i]
	}

	for i := 0; i < quantity; i++ {
		e := objects[i]
		l.hash[i+index] = &e
	}

	l.count += quantity
	return true
}

//Return true when the List is empty.
//O(1)
func (l *List) IsEmpty() bool {
	return l.count == 0
}

//Removes the first occurrence of a specific object from the List.
//O(n)
func (l *List) Remove(object T) bool {
	if l.IsEmpty() {
		return true
	}
	found := false
	for i := 0; i < l.count; i++ {
		if found {
			l.hash[i-1] = l.hash[i]
			continue
		}
		if *l.hash[i] == object {
			found = true
			continue
		}
	}
	l.count--
	delete(l.hash, l.count)
	return found
}

//Removes all the elements that match the conditions defined by the specified predicate.
//O(n)
func (l *List) RemoveAll(object T) bool {
	if l.IsEmpty() {
		return true
	}
	found := 0
	for i := 0; i < l.count; i++ {
		if *l.hash[i] == object {
			found++
			continue
		}
		if found != 0 {
			l.hash[i-found] = l.hash[i]
			continue
		}
	}
	l.count -= found
	for i := 0; i < found; i++ {
		delete(l.hash, l.count+i)
	}
	return found != 0
}

//Removes the element at the specified index of the List.
//O(n)
func (l *List) RemoveAt(index int) bool {
	if index < 0 || index >= l.count {
		return false
	}
	for i := index + 1; i < l.count; i++ {
		l.hash[i-1] = l.hash[i]
	}
	l.count--
	delete(l.hash, l.count)
	return true
}

//Sets the element at the specified index.
//Return nil when the index of List is empty.
//O(1)
func (l *List) Set(index int, object T) bool {
	if index < 0 || index > l.count || l.IsEmpty() {
		return false
	}

	l.hash[index] = &object
	return true
}

//Copies the elements of the List to a new array.
//O(n)
func (l *List) ToArray() []T {
	if l.IsEmpty() {
		return nil
	}

	arr := make([]T, l.count)
	for i := 0; i < l.count; i++ {
		arr[i] = *l.hash[i]
	}
	return arr
}
