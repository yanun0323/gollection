package gollection

import "sync"

type SyncBTree[T any] BTree[T]

func NewSyncBTree[T any](greater func(t1, t2 T) bool) SyncBTree[T] {
	return &syncBTree[T]{
		rwLock: &sync.RWMutex{},
		bt:     NewBTree[T](greater),
	}
}

type syncBTree[T any] struct {
	rwLock *sync.RWMutex
	bt     BTree[T]
}

func (b *syncBTree[T]) Len() int {
	b.rwLock.RLock()
	defer b.rwLock.RUnlock()
	return b.bt.Len()
}

func (b *syncBTree[T]) Insert(v T) {
	b.rwLock.Lock()
	defer b.rwLock.Unlock()
	b.bt.Insert(v)
}

func (b *syncBTree[T]) Walk(t int, limit ...int) []T {
	b.rwLock.RLock()
	defer b.rwLock.RUnlock()
	return b.bt.Walk(t, limit...)
}
