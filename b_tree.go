package gollection

// BTree is an interface for a binary tree data structure.
type BTree[K Orderable, V any] interface {
	// Contain returns true if the tree contains the key.
	Contain(key K) bool

	// Len returns the number of elements in the tree.
	Len() int

	// Insert adds a key-value pair to the tree.
	Insert(key K, value V)

	// Remove removes the key-value pair from the tree.
	Remove(key K) (V, bool)

	// Search returns the value associated with the key, or nil if the key is not found.
	Search(key K) (V, bool)

	// Max returns the key and value of the maximum element in the tree, or nil if the tree is empty.
	Max() (K, V, bool)

	// Min returns the key and value of the minimum element in the tree, or nil if the tree is empty.
	Min() (K, V, bool)

	// RemoveMax removes the key and value of the maximum element in the tree, or nil if the tree is empty.
	RemoveMax() (K, V, bool)

	// RemoveMin removes the key and value of the minimum element in the tree, or nil if the tree is empty.
	RemoveMin() (K, V, bool)

	// Ascend calls the given function for each key and value in the tree in ascending order.
	// The function should return true to continue the iteration or false to stop it.
	Ascend(fn TreeIter[K, V])

	// Descend calls the given function for each key and value in the tree in descending order.
	// The function should return true to continue the iteration or false to stop it.
	Descend(fn TreeIter[K, V])
}

type bTree[K Orderable, V any] struct {
	kZero K
	vZero V
	count int
	root  *node[K, V]
}

// NewBTree returns a new binary tree.
func NewBTree[K Orderable, V any](elems ...map[K]V) BTree[K, V] {
	t := &bTree[K, V]{}
	for _, e := range elems {
		for k, v := range e {
			t.Insert(k, v)
		}
	}

	return t
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

func (b *bTree[K, V]) ascend(n *node[K, V], fn TreeIter[K, V]) bool {
	if n == nil {
		return true
	}

	if !b.ascend(n.l, fn) {
		return false
	}

	if !fn(n.key, n.val) {
		return false
	}

	if !b.ascend(n.r, fn) {
		return false
	}

	return true
}

func (b *bTree[K, V]) Descend(fn TreeIter[K, V]) {
	b.descend(b.root, fn)
}

func (b *bTree[K, V]) descend(n *node[K, V], fn TreeIter[K, V]) bool {
	if n == nil {
		return true
	}
	if !b.descend(n.r, fn) {
		return false
	}

	if !fn(n.key, n.val) {
		return false
	}

	if !b.descend(n.l, fn) {
		return false
	}

	return true
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
