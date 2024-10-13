package gollection

import "cmp"

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
//
// Note that floating-point types may contain NaN ("not-a-number") values.
// An operator such as == or < will always report false when
// comparing a NaN value with any other value, NaN or not.
// See the [Compare] function for a consistent way to compare NaN values.
type Orderable cmp.Ordered

// Compare is a function that compares two values and returns true if the first value is less than the second value.
type Compare[T any] func(T, T) bool

type Iter[T any] func(T) bool

// MapIter is a function that iterates over a map.
// The function should return true to continue the iteration or false to stop it.
type MapIter[K comparable, V any] func(K, V) bool

// TreeIter is a function that iterates over a tree.
// The function should return true to continue the iteration or false to stop it.
type TreeIter[K Orderable, V any] func(K, V) bool
