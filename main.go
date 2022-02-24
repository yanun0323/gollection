package main

import (
	"main/datastruct"
)

func main() {
	//	others.ValueOrPointer()
	//datastruct.Test_stack()
	s := datastruct.NewStack()
	q := datastruct.NewQueue()
	s.Push(nil)
	q.Enqueue(nil)
}
