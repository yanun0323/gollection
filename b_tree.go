package gollection

type BTree[K orderable, V any] interface {
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

func NewBTree[K orderable, V any]() BTree[K, V] {
	return &bTree[K, V]{}
}

type bTree[K orderable, V any] struct {
	kZero K
	vZero V
	count int
	root  *node[K, V]
}

func (b *bTree[K, V]) Contain(key K) bool {
	if b.root == nil {
		return false
	}
	return b.root.Find(key) != nil
}

func (b *bTree[K, V]) Len() int {
	return b.count
}

func (b *bTree[K, V]) Insert(key K, value V) {
	n := newNode[K, V](key, value)
	b.count++
	if b.root == nil {
		b.root = n
		return
	}
	b.root.Insert(n)
}

func (b *bTree[K, V]) Remove(key K) (V, bool) {
	if b.root == nil {
		return b.vZero, false
	}
	h, removed := b.root.FindAndRemove(key)
	if removed == nil {
		return b.vZero, false
	}
	b.root = h
	b.count--
	return removed.val, true
}

func (b *bTree[K, V]) Search(key K) (V, bool) {
	if b.root == nil {
		return b.vZero, false
	}
	n := b.root.Find(key)
	if n == nil {
		return b.vZero, false
	}
	return n.val, true
}

func (b *bTree[K, V]) Max() (K, V, bool) {
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	n := b.root.FindMax()
	if n == nil {
		return b.kZero, b.vZero, false
	}
	return n.key, n.val, true
}

func (b *bTree[K, V]) Min() (K, V, bool) {
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	n := b.root.FindMin()
	if n == nil {
		return b.kZero, b.vZero, false
	}
	return n.key, n.val, true
}

func (b *bTree[K, V]) RemoveMax() (K, V, bool) {
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	if b.root.r == nil {
		result := b.root
		b.root = b.root.l
		b.count--
		return result.key, result.val, true
	}

	result, c := b.root.remove(_max)
	b.root.r = c
	if result == nil {
		return b.kZero, b.vZero, false
	}
	b.count--
	return result.key, result.val, true
}

func (b *bTree[K, V]) RemoveMin() (K, V, bool) {
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	if b.root.l == nil {
		result := b.root
		b.root = b.root.r
		b.count--
		return result.key, result.val, true
	}

	result, c := b.root.remove(_min)
	b.root.l = c
	if result == nil {
		return b.kZero, b.vZero, false
	}
	b.count--
	return result.key, result.val, true
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

type node[K orderable, V any] struct {
	key K
	val V
	l   *node[K, V]
	r   *node[K, V]
}

func newNode[K orderable, V any](key K, value V) *node[K, V] {
	return &node[K, V]{
		key: key,
		val: value,
	}
}

func (n *node[K, V]) Insert(nn *node[K, V]) {
	if nn == nil {
		return
	}
	if nn.key <= n.key {
		if n.l == nil {
			n.l = nn
			return
		}
		n.l.Insert(nn)
	} else {
		if n.r == nil {
			n.r = nn
			return
		}
		n.r.Insert(nn)
	}
}

func (n *node[K, V]) Find(key K) *node[K, V] {
	if n.key == key {
		return n
	}

	if n.key < n.key {
		if n.l == nil {
			return nil
		}
		return n.l.Find(key)
	} else {
		if n.r == nil {
			return nil
		}
		return n.r.Find(key)
	}
}

func (n *node[K, V]) FindMax() *node[K, V] {
	if n.r != nil {
		return n.r.FindMax()
	}
	return n
}

func (n *node[K, V]) FindMin() *node[K, V] {
	if n.l != nil {
		return n.l.FindMax()
	}
	return n
}

type side bool

const (
	_max side = false
	_min side = true
)

func (n *node[K, V]) FindAndRemove(key K) (heir, removed *node[K, V]) {
	if key > n.key { // find into right to remove
		if n.r == nil {
			return n, nil
		}
		c, removed := n.r.FindAndRemove(key)
		n.r = c
		return n, removed
	}

	if key < n.key { // find into left to remove
		if n.l == nil {
			return n, nil
		}
		c, removed := n.l.FindAndRemove(key)
		n.l = c
		return n, removed
	}

	// key == n.key: remove current node
	left, right := n.removeChildren()
	if left == nil {
		return right, n
	}

	if right == nil {
		return left, n
	}

	heir, c := left.remove(_max)
	heir.r = right
	heir.l = c
	return heir, n
}

func (n *node[K, V]) removeChildren() (left, right *node[K, V]) {
	l, r := n.l, n.r
	n.l, n.r = nil, nil
	return l, r
}

func (n *node[K, V]) remove(next side) (removed *node[K, V], child *node[K, V]) {
	switch next {
	case _max:
		if n.r != nil {
			nr, nc := n.r.remove(next)
			n.r = nc
			return nr, nil
		}
		left, _ := n.removeChildren()
		return n, left
	case _min:
		if n.l != nil {
			nr, nc := n.l.remove(next)
			n.l = nc
			return nr, nil
		}
		_, right := n.removeChildren()
		return n, right
	default:
		return n, nil
	}
}
