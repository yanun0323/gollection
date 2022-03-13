package gollection

type IQueue interface {
	Clear()
	Clone() IQueue
	Contains(...interface{}) bool
	Count() int
	Dequeue() interface{}
	Enqueue(interface{})
	IsEmpty() bool
	Peek() interface{}
	ToArray() []interface{}
}

type IStack interface {
	Clear()
	Clone() IStack
	Contains(...interface{}) bool
	Count() int
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
	Contains(...interface{}) bool
	Count() int
	Insert(int, ...interface{})
	IsEmpty() bool
	Remove(interface{})
	RemoveAll(interface{})
	RemoveAt(int) interface{}
	Set(int, interface{})
	ToArray() []interface{}
}
