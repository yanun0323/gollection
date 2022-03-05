package gollection

type IQueue interface {
	Clear()
	Clone() IQueue
	Contains(interface{}) bool
	ContainsAny(...interface{}) bool
	Dequeue() interface{}
	Enqueue(interface{})
	IsEmpty() bool
	Peek() interface{}
	ToArray() []interface{}
}

type IStack interface {
	Clear()
	Clone() IStack
	Contains(interface{}) bool
	ContainsAny(...interface{}) bool
	IsEmpty() bool
	Peek() interface{}
	Pop() interface{}
	Push(interface{})
	ToArray() []interface{}
}

type IList interface {
	ADD(...interface{})
	At(int) interface{}
	Clear()
	Clone() IList
	Contains(interface{}) bool
	ContainsAny(...interface{}) bool
	Insert(int, ...interface{}) bool
	IsEmpty() bool
	Remove(interface{}) bool
	RemoveAll(interface{}) bool
	RemoveAt(int) bool
	Set(int, interface{}) bool
	ToArray() []interface{}
}
