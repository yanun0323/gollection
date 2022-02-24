package main

import (
	"main/datastruct"
)

func main() {
	s := datastruct.NewStack()
	q := datastruct.NewQueue()
	s.Push(nil)
	q.Enqueue(nil)
}
