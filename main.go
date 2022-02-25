package main

import (
	"fmt"

	"github.com/yanun0323/gollection/datastruct"
)

func main() {
	q := datastruct.NewQueue()
	q.Enqueue("A")

	s1 := "Hi"
	fmt.Println(&s1)
	s2 := "Hi"
	fmt.Println(&s2)
	s2 = "hi"
	fmt.Println(&s2)
}
