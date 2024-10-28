package gollection

// AvlTree is an interface for an AVL tree data structure.
type AvlTree[K Orderable, V any] interface {
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

	// Clear removes all elements from the tree.
	Clear()
}

type avlTree[K Orderable, V any] struct {
	kZero K
	vZero V
	count int
	root  *avlNode[K, V]
}

// NewAvlTree returns a new AVL tree.
func NewAvlTree[K Orderable, V any](elems ...map[K]V) AvlTree[K, V] {
	t := &avlTree[K, V]{}

	for _, e := range elems {
		for k, v := range e {
			t.Insert(k, v)
		}
	}

	return t
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

func (a *avlTree[K, V]) ascend(n *avlNode[K, V], fn TreeIter[K, V]) bool {
	if n == nil {
		return true
	}

	if !a.ascend(n.l, fn) {
		return false
	}

	if !fn(n.key, n.val) {
		return false
	}

	if !a.ascend(n.r, fn) {
		return false
	}

	return true
}

func (a *avlTree[K, V]) Descend(fn TreeIter[K, V]) {
	a.descend(a.root, fn)
}

func (a *avlTree[K, V]) descend(n *avlNode[K, V], fn TreeIter[K, V]) bool {
	if n == nil {
		return true
	}

	if !a.descend(n.r, fn) {
		return false
	}

	if !fn(n.key, n.val) {
		return false
	}

	if !a.descend(n.l, fn) {
		return false
	}

	return true
}

type avlNode[K Orderable, V any] struct {
	key    K
	val    V
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

// rebalancedTree checks if node is balanced and rebalanced
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

// rotateLeft rotates nodes left to balance node
func (n *avlNode[K, V]) rotateLeft() *avlNode[K, V] {
	newRoot := n.r
	n.r = newRoot.l
	newRoot.l = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// rotateRight rotates nodes right to balance node
func (n *avlNode[K, V]) rotateRight() *avlNode[K, V] {
	newRoot := n.l
	n.l = newRoot.r
	newRoot.r = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (a *avlTree[K, V]) Clear() {
	a.root = nil
	a.count = 0
}
