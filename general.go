package gollection

import "cmp"

type Orderable cmp.Ordered

type Compare[T any] func(T, T) bool

type Iter[T any] func(T) bool

type MapIter[K comparable, V any] func(K, V) bool

type TreeIter[K Orderable, V any] func(K, V) bool
