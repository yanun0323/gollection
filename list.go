package gollection

import (
	"github.com/yanun0323/gollection/errors"
)

//Represents a strongly typed list of objects that can be accessed by index. Provides methods to search, sort, and manipulate lists.
//Max stored element quantity is 2,147,483,647
type List struct {
	count int
	hash  map[int]*interface{}
}

//Initializes a new instance of the List class that is empty and has the default initial capacity.
//
//O(1)
func NewList(objects ...interface{}) List {
	l := List{count: 0, hash: map[int]*interface{}{}}
	if len(objects) == 0 {
		return l
	}
	l.ADD(objects...)
	return l
}

//Adds an object to the end of the List.
//
//O(1)
func (l *List) ADD(objects ...interface{}) {
	if len(objects) == 0 {
		return
	}
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		l.hash[l.count] = &obj
		l.count++
	}
}

//Gets the element at the specified index.
//Panic when the index of the List is empty.
//
//O(1)
func (l *List) At(index int) interface{} {
	if index < 0 || l.IsEmpty() {
		panic(errors.OutOfListBounds)
	}

	obj, ok := l.hash[index]
	if !ok {
		panic(errors.EmptyListIndex(index))
	}

	return *obj
}

//Removes all elements from the List.
//
//O(1)
func (l *List) Clear() {
	l.count = 0
	l.hash = map[int]*interface{}{}
}

//Clone the List without clone the objects inside the List.
//
//O(n)
func (l *List) Clone() List {
	hash := map[int]*interface{}{}
	for key, value := range l.hash {
		hash[key] = value
	}
	return List{count: l.count, hash: hash}
}

//Determines whether any element is in the List.
//
//O(n)
func (l *List) Contains(objects ...interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	if len(objects) == 0 {
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
//
//O(1)
func (l *List) Count() int {
	return l.count
}

//Inserts any element into the List at the specified index.
//Panic when the index of the List is empty.
//
//O(n)
func (l *List) Insert(index int, objects ...interface{}) {
	if index < 0 || index > l.count {
		panic(errors.OutOfListBounds)
	}
	if len(objects) == 0 {
		return
	}
	if index == l.count {
		l.ADD(objects...)
		return
	}

	quantity := len(objects)
	for i := index; i < l.count; i++ {
		l.hash[i+quantity] = l.hash[i]
	}

	for i := 0; i < quantity; i++ {
		obj := objects[i]
		l.hash[i+index] = &obj
	}

	l.count += quantity
}

//Return true when the List is empty.
//
//O(1)
func (l *List) IsEmpty() bool {
	return l.count == 0
}

//Removes the first object at the List.
//
//O(n)
func (l *List) Remove(object interface{}) {
	if l.IsEmpty() {
		return
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

	if !found {
		return
	}

	l.count--
	delete(l.hash, l.count)
}

//Removes all the elements that match the conditions defined by the specified predicate.
//
//O(n)
func (l *List) RemoveAll(object interface{}) {
	if l.IsEmpty() {
		return
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

	if found == 0 {
		return
	}

	l.count -= found
	for i := 0; i < found; i++ {
		delete(l.hash, l.count+i)
	}
}

//Removes the element at the specified index of the List.
//Panic when the index of the List is empty.
//
//O(n)
func (l *List) RemoveAt(index int) interface{} {
	if index < 0 || index >= l.count {
		panic(errors.OutOfListBounds)
	}
	d := *l.hash[index]
	for i := index + 1; i < l.count; i++ {
		l.hash[i-1] = l.hash[i]
	}
	l.count--
	delete(l.hash, l.count)
	return d
}

//Sets the element at the specified index.
//Panic when the index of List is empty.
//
//O(1)
func (l *List) Set(index int, object interface{}) {
	if index < 0 || index > l.count || l.IsEmpty() {
		panic(errors.OutOfListBounds)
	}

	l.hash[index] = &object
}

//Copies the elements of the List to a new slice.
//Return nil when the List is empty.
//
//O(n)
func (l *List) ToArray() []interface{} {
	if l.IsEmpty() {
		return nil
	}

	arr := make([]interface{}, l.count)
	for i := 0; i < l.count; i++ {
		arr[i] = *l.hash[i]
	}
	return arr
}
