package main

import (
	"fmt"
	"main/gollection"
)

func main() {
	l := gollection.NewList()
	l.ADD(1, 0, 2, 3, 0, 4, 5, 6, 0, 7)
	// for i := 0; i < l.Count; i++ {
	// 	fmt.Println(l.At(i).(int))
	// }
	ok := l.RemoveAll(0)
	l.RemoveAt(7)
	fmt.Println(ok)

	for _, e := range l.ToArray() {
		fmt.Println(e.(int))
	}
}
