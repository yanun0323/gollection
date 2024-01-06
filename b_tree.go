package gollection

type BTree[K Orderable, V any] interface {
	Contain(key K) bool
	Len() int
	Insert(key K, value V)
	Remove(key K) (V, bool)
	Search(key K) (V, bool)
	Max() (K, V, bool)
	Min() (K, V, bool)
	RemoveMax() (K, V, bool)
	RemoveMin() (K, V, bool)
	Ascend(fn TreeIter[K, V])
	Descend(fn TreeIter[K, V])
}

func NewBTree[K Orderable, V any]() BTree[K, V] {
	return &bTree[K, V]{}
}

type bTree[K Orderable, V any] struct {
	kZero K
	vZero V
	count int
	root  *node[K, V]
}

func (b *bTree[K, V]) Contain(key K) bool {
	return b.root.search(key) != nil
}

func (b *bTree[K, V]) Len() int {
	return b.count
}

func (b *bTree[K, V]) Insert(key K, value V) {
	new := false
	b.root = b.root.add(key, value, &new)
	if new {
		b.count++
	}
}

func (b *bTree[K, V]) Remove(key K) (V, bool) {
	var (
		removed node[K, V]
		ok      bool
	)
	b.root = b.root.remove(key, &removed, &ok)
	if ok {
		b.count--
	}
	return removed.val, ok
}

func (b *bTree[K, V]) Search(key K) (V, bool) {
	if n := b.root.search(key); n != nil {
		return n.val, true
	}
	return b.vZero, false
}

func (b *bTree[K, V]) Max() (K, V, bool) {
	if n := b.root.findMax(); n != nil {
		return n.key, n.val, true
	}
	return b.kZero, b.vZero, false
}

func (b *bTree[K, V]) Min() (K, V, bool) {
	if n := b.root.findMin(); n != nil {
		return n.key, n.val, true
	}
	return b.kZero, b.vZero, false
}

func (b *bTree[K, V]) RemoveMax() (K, V, bool) {
	var (
		removed node[K, V]
		ok      bool
	)
	b.root = b.root.removeMax(&removed, &ok)
	if ok {
		b.count--
	}
	return removed.key, removed.val, ok
}

func (b *bTree[K, V]) RemoveMin() (K, V, bool) {
	var (
		removed node[K, V]
		ok      bool
	)
	b.root = b.root.removeMin(&removed, &ok)
	if ok {
		b.count--
	}
	return removed.key, removed.val, ok
}

func (b *bTree[K, V]) Ascend(fn TreeIter[K, V]) {
	b.ascend(b.root, fn)
}

func (b *bTree[K, V]) ascend(n *node[K, V], fn TreeIter[K, V]) {
	if n == nil {
		return
	}
	b.ascend(n.l, fn)
	fn(n.key, n.val)
	b.ascend(n.r, fn)
}

func (b *bTree[K, V]) Descend(fn TreeIter[K, V]) {
	b.descend(b.root, fn)
}

func (b *bTree[K, V]) descend(n *node[K, V], fn TreeIter[K, V]) {
	if n == nil {
		return
	}
	b.descend(n.r, fn)
	fn(n.key, n.val)
	b.descend(n.l, fn)
}

type node[K Orderable, V any] struct {
	key K
	val V
	l   *node[K, V]
	r   *node[K, V]
}

// add adds or updates node value
func (n *node[K, V]) add(key K, val V, new *bool) *node[K, V] {
	if n == nil {
		*new = true
		return &node[K, V]{key, val, nil, nil}
	}

	if key < n.key {
		n.l = n.l.add(key, val, new)
	} else if key > n.key {
		n.r = n.r.add(key, val, new)
	} else {
		n.val = val
	}
	return n
}

func (n *node[K, V]) search(key K) *node[K, V] {
	if n == nil {
		return n
	}

	if key < n.key {
		return n.l.search(key)
	}

	if key > n.key {
		return n.r.search(key)
	}

	return n
}

func (n *node[K, V]) findMax() *node[K, V] {
	if n == nil {
		return nil
	}

	if r := n.r.findMax(); r != nil {
		return r
	}

	return n
}

func (n *node[K, V]) findMin() *node[K, V] {
	if n == nil {
		return nil
	}

	if l := n.l.findMin(); l != nil {
		return l
	}
	return n
}

func (n *node[K, V]) remove(key K, removed *node[K, V], ok *bool) *node[K, V] {
	if n == nil {
		return nil
	}

	if key > n.key {
		n.r = n.r.remove(key, removed, ok)
		return n
	}

	if key < n.key {
		n.l = n.l.remove(key, removed, ok)
		return n
	}

	*removed = *n
	*ok = true
	defer func() {
		n.l, n.r = nil, nil
		n = nil
	}()

	if n.l == nil {
		return n.r
	}

	if n.r == nil {
		return n.l
	}

	var (
		rightMinNode node[K, V]
		deleted      bool
	)
	n.r = n.r.removeMin(&rightMinNode, &deleted)
	rightMinNode.r = n.r
	rightMinNode.l = n.l
	return &rightMinNode
}

func (n *node[K, V]) removeMax(removed *node[K, V], ok *bool) *node[K, V] {
	if n == nil {
		return nil
	}

	if n.r == nil {
		defer func() {
			n.r, n.l = nil, nil
			n = nil
		}()
		*ok = true
		*removed = *n
		return n.l
	}

	n.r = n.r.removeMax(removed, ok)
	return n
}

func (n *node[K, V]) removeMin(removed *node[K, V], ok *bool) *node[K, V] {
	if n == nil {
		return nil
	}

	if n.l == nil {
		defer func() {
			n.r, n.l = nil, nil
			n = nil
		}()
		*ok = true
		*removed = *n
		return n.r
	}

	n.l = n.l.removeMin(removed, ok)
	return n
}
