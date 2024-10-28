package gollection

import "sync"

type syncAvlTree[K Orderable, V any] struct {
	rwLock *sync.RWMutex
	t      BTree[K, V]
}

// NewSyncAvlTree returns a new thread-safe AVL tree.
func NewSyncAvlTree[K Orderable, V any](elems ...map[K]V) AvlTree[K, V] {
	return &syncAvlTree[K, V]{
		rwLock: &sync.RWMutex{},
		t:      NewAvlTree[K, V](elems...),
	}
}

func (a *syncAvlTree[K, V]) Contain(key K) bool {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	return a.t.Contain(key)
}

func (a *syncAvlTree[K, V]) Len() int {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	return a.t.Len()
}

func (a *syncAvlTree[K, V]) Insert(key K, value V) {
	a.rwLock.Lock()
	defer a.rwLock.Unlock()
	a.t.Insert(key, value)
}

func (a *syncAvlTree[K, V]) Remove(key K) (V, bool) {
	a.rwLock.Lock()
	defer a.rwLock.Unlock()
	return a.t.Remove(key)
}

func (a *syncAvlTree[K, V]) Search(key K) (V, bool) {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	return a.t.Search(key)
}

func (a *syncAvlTree[K, V]) Max() (K, V, bool) {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	return a.t.Max()
}

func (a *syncAvlTree[K, V]) Min() (K, V, bool) {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	return a.t.Min()
}

func (a *syncAvlTree[K, V]) RemoveMax() (K, V, bool) {
	a.rwLock.Lock()
	defer a.rwLock.Unlock()
	return a.t.RemoveMax()
}

func (a *syncAvlTree[K, V]) RemoveMin() (K, V, bool) {
	a.rwLock.Lock()
	defer a.rwLock.Unlock()
	return a.t.RemoveMin()
}

func (a *syncAvlTree[K, V]) Ascend(fn TreeIter[K, V]) {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	a.t.Ascend(fn)
}

func (a *syncAvlTree[K, V]) Descend(fn TreeIter[K, V]) {
	a.rwLock.RLock()
	defer a.rwLock.RUnlock()
	a.t.Descend(fn)
}

func (a *syncAvlTree[K, V]) Clear() {
	a.rwLock.Lock()
	defer a.rwLock.Unlock()
	a.t.Clear()
}
