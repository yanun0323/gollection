package gollection

import "cmp"

type orderable cmp.Ordered

type Compare[T any] func(T, T) bool

type Iter[T any] func(T) bool

type MapIter[K comparable, V any] func(K, V) bool

type TreeIter[K orderable, V any] func(K, V) bool
