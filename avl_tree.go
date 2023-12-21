package gollection

type AVLTree[K orderable, V any] interface {
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

func NewAVLTree[K orderable, V any]() AVLTree[K, V] {
	return &avlTree[K, V]{}
}

// avlTree[K,V] structure. Public methods are Add, Remove, Update, Search, DisplayTreeInOrder.
type avlTree[K orderable, V any] struct {
	kZero K
	vZero V
	count int
	root  *avlNode[K, V]
}

func (a *avlTree[K, V]) Contain(key K) bool {
	return a.root.search(key) != nil
}

func (a *avlTree[K, V]) Len() int {
	return a.count
}

func (a *avlTree[K, V]) Insert(key K, value V) {
	new := false
	a.root = a.root.add(key, value, &new)
	if new {
		a.count++
	}
}

func (a *avlTree[K, V]) Remove(key K) (V, bool) {
	var (
		removed avlNode[K, V]
		ok      bool
	)
	a.count--
	a.root = a.root.remove(key, &removed, &ok)
	if ok {
		a.count--
	}
	return removed.val, ok
}

func (a *avlTree[K, V]) Search(key K) (V, bool) {
	if n := a.root.search(key); n != nil {
		return n.val, true
	}
	return a.vZero, false
}

func (a *avlTree[K, V]) Max() (K, V, bool) {
	if n := a.root.findMax(); n != nil {
		return n.key, n.val, true
	}
	return a.kZero, a.vZero, false
}

func (a *avlTree[K, V]) Min() (K, V, bool) {
	if n := a.root.findMin(); n != nil {
		return n.key, n.val, true
	}
	return a.kZero, a.vZero, false
}

func (a *avlTree[K, V]) RemoveMax() (K, V, bool) {
	var (
		removed avlNode[K, V]
		ok      bool
	)
	a.root = a.root.removeMax(&removed, &ok)
	if ok {
		a.count--
	}
	return removed.key, removed.val, ok
}

func (a *avlTree[K, V]) RemoveMin() (K, V, bool) {
	var (
		removed avlNode[K, V]
		ok      bool
	)
	a.root = a.root.removeMin(&removed, &ok)
	if ok {
		a.count--
	}
	return removed.key, removed.val, ok
}

func (a *avlTree[K, V]) Ascend(fn TreeIter[K, V]) {
	a.ascend(a.root, fn)
}

func (a *avlTree[K, V]) ascend(n *avlNode[K, V], fn TreeIter[K, V]) {
	if n == nil {
		return
	}
	a.ascend(n.l, fn)
	fn(n.key, n.val)
	a.ascend(n.r, fn)
}

func (a *avlTree[K, V]) Descend(fn TreeIter[K, V]) {
	a.descend(a.root, fn)
}

func (a *avlTree[K, V]) descend(n *avlNode[K, V], fn TreeIter[K, V]) {
	if n == nil {
		return
	}
	a.descend(n.r, fn)
	fn(n.key, n.val)
	a.descend(n.l, fn)
}

// avlNode structure
type avlNode[K orderable, V any] struct {
	key K
	val V

	// height counts nodes (not edges)
	height int
	l      *avlNode[K, V]
	r      *avlNode[K, V]
}

// add adds or updates node value
func (n *avlNode[K, V]) add(key K, value V, new *bool) *avlNode[K, V] {
	if n == nil {
		*new = true
		return &avlNode[K, V]{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.l = n.l.add(key, value, new)
	} else if key > n.key {
		n.r = n.r.add(key, value, new)
	} else {
		// if same key exists update value
		n.val = value
	}
	return n.rebalancedTree()
}

// Searches for a node
func (n *avlNode[K, V]) search(key K) *avlNode[K, V] {
	if n == nil {
		return nil
	}

	if key < n.key {
		return n.l.search(key)
	}

	if key > n.key {
		return n.r.search(key)
	}

	return n
}

func (n *avlNode[K, V]) findMax() *avlNode[K, V] {
	if n == nil {
		return nil
	}

	if r := n.r.findMax(); r != nil {
		return r
	}

	return n
}

func (n *avlNode[K, V]) findMin() *avlNode[K, V] {
	if n == nil {
		return nil
	}

	if l := n.l.findMin(); l != nil {
		return l
	}
	return n
}

// Removes a node
func (n *avlNode[K, V]) remove(key K, removed *avlNode[K, V], ok *bool) *avlNode[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.l = n.l.remove(key, removed, ok)
		return n.rebalancedTree()
	}
	if key > n.key {
		n.r = n.r.remove(key, removed, ok)
		return n.rebalancedTree()
	}

	{
		*removed = *n
		*ok = true
		defer func() {
			n.r, n.l = nil, nil
		}()

		if n.l == nil {
			return n.r.rebalancedTree()
		}

		if n.r == nil {
			return n.l.rebalancedTree()
		}

		var (
			rightMinNode avlNode[K, V]
			deleted      bool
		)

		n.r = n.r.removeMin(&rightMinNode, &deleted)
		rightMinNode.r = n.r
		rightMinNode.l = n.l
		return rightMinNode.rebalancedTree()
	}
}

func (n *avlNode[K, V]) removeMax(removed *avlNode[K, V], ok *bool) *avlNode[K, V] {
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
		return n.l.rebalancedTree()
	}

	n.r = n.r.removeMax(removed, ok)
	return n.rebalancedTree()
}

func (n *avlNode[K, V]) removeMin(removed *avlNode[K, V], ok *bool) *avlNode[K, V] {
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
		return n.r.rebalancedTree()
	}

	n.l = n.l.removeMin(removed, ok)
	return n
}

func (n *avlNode[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode[K, V]) recalculateHeight() {
	n.height = 1 + max(n.l.getHeight(), n.r.getHeight())
}

// Checks if node is balanced and rebalanced
func (n *avlNode[K, V]) rebalancedTree() *avlNode[K, V] {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.l.getHeight() - n.r.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.r.l.getHeight() > n.r.r.getHeight() {
			n.r = n.r.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.l.r.getHeight() > n.l.l.getHeight() {
			n.l = n.l.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *avlNode[K, V]) rotateLeft() *avlNode[K, V] {
	newRoot := n.r
	n.r = newRoot.l
	newRoot.l = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *avlNode[K, V]) rotateRight() *avlNode[K, V] {
	newRoot := n.l
	n.l = newRoot.r
	newRoot.r = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}
