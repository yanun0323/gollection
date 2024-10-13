package gollection

import "cmp"

type orderable cmp.Ordered

// Compare is a function that compares two values and returns true if the first value is less than the second value.
type Compare[T any] func(T, T) bool

type Iter[T any] func(T) bool

// MapIter is a function that iterates over a map.
// The function should return true to continue the iteration or false to stop it.
type MapIter[K comparable, V any] func(K, V) bool

// TreeIter is a function that iterates over a tree.
// The function should return true to continue the iteration or false to stop it.
type TreeIter[K orderable, V any] func(K, V) bool
