package gollection

type T interface{}

type IQueue interface {
	Clear()
	Clone() IQueue
	Contains(T) bool
	ContainsAny(...T) bool
	Count() int
	Dequeue() T
	Enqueue(T)
	IsEmpty() bool
	Peek() T
	ToArray() []T
}

type IStack interface {
	Clear()
	Clone() IStack
	Contains(T) bool
	ContainsAny(...T) bool
	Count() int
	IsEmpty() bool
	Peek() T
	Pop() T
	Push(T)
	ToArray() []T
}

type IList interface {
	ADD(...T)
	At(int) T
	Clear()
	Clone() IList
	Contains(T) bool
	Count() int
	ContainsAny(...T) bool
	Insert(int, ...T) bool
	IsEmpty() bool
	Remove(T) bool
	RemoveAll(T) bool
	RemoveAt(int) bool
	Set(int, T) bool
	ToArray() []T
}
