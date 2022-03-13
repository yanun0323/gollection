package errors

import "fmt"

var (
	EmptyQueue      = "No element in queue"
	EmptyStack      = "No element in stack"
	EmptyList       = "No element in list"
	OutOfListBounds = "Out of bounds"
)

func EmptyListIndex(i int) string {
	return fmt.Sprintf("No element at %d", i)
}
