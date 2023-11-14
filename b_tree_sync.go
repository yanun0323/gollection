package gollection

import "sync"

// import "sync"

type SyncBTree[K orderable, V any] BTree[K, V]

func NewSyncBTree[K orderable, V any]() SyncBTree[K, V] {
	return &syncBTree[K, V]{
		rwLock: &sync.RWMutex{},
		bt:     NewBTree[K, V](),
	}
}

type syncBTree[K orderable, V any] struct {
	rwLock *sync.RWMutex
	bt     BTree[K, V]
}

func (s *syncBTree[K, V]) Contain(key K) bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.bt.Contain(key)
}

func (s *syncBTree[K, V]) Len() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.bt.Len()
}

func (s *syncBTree[K, V]) Insert(key K, value V) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.bt.Insert(key, value)
}

func (s *syncBTree[K, V]) Remove(key K) (V, bool) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.bt.Remove(key)
}

func (s *syncBTree[K, V]) Search(key K) (V, bool) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.bt.Search(key)
}

func (s *syncBTree[K, V]) Max() (K, V, bool) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.bt.Max()
}

func (s *syncBTree[K, V]) Min() (K, V, bool) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.bt.Min()
}

func (s *syncBTree[K, V]) RemoveMax() (K, V, bool) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.bt.RemoveMax()
}

func (s *syncBTree[K, V]) RemoveMin() (K, V, bool) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.bt.RemoveMin()
}

func (s *syncBTree[K, V]) Ascend(fn TreeIter[K, V]) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	s.bt.Ascend(fn)
}

func (s *syncBTree[K, V]) Descend(fn TreeIter[K, V]) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	s.bt.Descend(fn)
}
