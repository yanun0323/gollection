package gollection_test

import (
	"fmt"

	"github.com/yanun0323/gollection/v2"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := gollection.NewLinkedList[int]()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
