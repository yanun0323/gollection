package gollection

type IQueue interface {
	Clear() bool
	Clone() IQueue
	Contains(...interface{}) bool
	Count() int
	Dequeue() interface{}
	Enqueue(interface{}) bool
	IsEmpty() bool
	Peek() interface{}
	ToArray() []interface{}
}

type IStack interface {
	Clear() bool
	Clone() IStack
	Contains(...interface{}) bool
	Count() int
	IsEmpty() bool
	Peek() interface{}
	Pop() interface{}
	Push(interface{}) bool
	ToArray() []interface{}
}

type IList interface {
	ADD(...interface{}) bool
	At(int) interface{}
	Clear() bool
	Clone() IList
	Contains(...interface{}) bool
	Count() int
	Insert(int, ...interface{}) bool
	IsEmpty() bool
	Remove(interface{}) bool
	RemoveAll(interface{}) bool
	RemoveAt(int) interface{}
	Set(int, interface{}) bool
	ToArray() []interface{}
}
