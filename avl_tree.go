package gollection

import (
	"fmt"
)

type AVLTree[K orderable, V any] interface {
	Contain(key K) bool
	Len() int
	Add(key K, value V)
	DisplayInOrder()
	Remove(key K)
	Search(key K) (V, bool)
	// Update(oldKey K, newKey K, newValue V)
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

func (t *avlTree[K, V]) Contain(key K) bool {
	return t.root.search(key) != nil
}

func (t *avlTree[K, V]) Len() int {
	return t.count
}

func (t *avlTree[K, V]) Add(key K, value V) {
	t.count++
	t.root = t.root.add(key, value)
}

func (t *avlTree[K, V]) Remove(key K) {
	t.count--
	t.root = t.root.remove(key)
}

func (t *avlTree[K, V]) Search(key K) (V, bool) {
	n := t.root.search(key)
	if n == nil {
		return t.vZero, false
	}
	return n.Value, true
}

func (t *avlTree[K, V]) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

// avlNode structure
type avlNode[K orderable, V any] struct {
	key   K
	Value V

	// height counts nodes (not edges)
	height int
	left   *avlNode[K, V]
	right  *avlNode[K, V]
}

// Adds a new node
func (n *avlNode[K, V]) add(key K, value V) *avlNode[K, V] {
	if n == nil {
		return &avlNode[K, V]{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		// if same key exists update value
		n.Value = value
	}
	return n.rebalanceTree()
}

// Removes a node
func (n *avlNode[K, V]) remove(key K) *avlNode[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			// delete smallest node that we replaced
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

// Searches for a node
func (n *avlNode[K, V]) search(key K) *avlNode[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

// Displays nodes left-depth first (used for debugging)
func (n *avlNode[K, V]) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *avlNode[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode[K, V]) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// Checks if node is balanced and rebalance
func (n *avlNode[K, V]) rebalanceTree() *avlNode[K, V] {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *avlNode[K, V]) rotateLeft() *avlNode[K, V] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *avlNode[K, V]) rotateRight() *avlNode[K, V] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Finds the smallest child (based on the key) for the current node
func (n *avlNode[K, V]) findSmallest() *avlNode[K, V] {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

// Returns max number - TODO: std lib seemed to only have a method for floats!
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
