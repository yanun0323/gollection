package gollection

type IQueue interface {
	Clear() bool
	Clone() IQueue
	Contains(...interface{}) bool
	Count() int
	Dequeue() (interface{}, bool)
	Enqueue(interface{}) bool
	IsEmpty() bool
	Peek() (interface{}, bool)
	ToArray() ([]interface{}, bool)
}

type IStack interface {
	Clear() bool
	Clone() IStack
	Contains(...interface{}) bool
	Count() int
	IsEmpty() bool
	Peek() (interface{}, bool)
	Pop() (interface{}, bool)
	Push(interface{}) bool
	ToArray() ([]interface{}, bool)
}

type IList interface {
	ADD(...interface{}) bool
	At(int) (interface{}, bool)
	Clear() bool
	Clone() IList
	Contains(...interface{}) bool
	Count() int
	Insert(int, ...interface{}) bool
	IsEmpty() bool
	Remove(interface{}) bool
	RemoveAll(interface{}) bool
	RemoveAt(int) (interface{}, bool)
	Set(int, interface{}) bool
	ToArray() ([]interface{}, bool)
}
