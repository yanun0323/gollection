package gollection

import (
	"strings"
	"testing"
)

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
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	n := b.root.findMax()
	if n == nil {
		return b.kZero, b.vZero, false
	}
	return n.key, n.val, true
}

func (b *bTree[K, V]) Min() (K, V, bool) {
	if b.root == nil {
		return b.kZero, b.vZero, false
	}
	n := b.root.findMin()
	if n == nil {
		return b.kZero, b.vZero, false
	}
	return n.key, n.val, true
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

type node[K orderable, V any] struct {
	key K
	val V
	l   *node[K, V]
	r   *node[K, V]
}

func newNode[K orderable, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		key: key,
		val: val,
	}
}

// add adds or updates node value
func (n *node[K, V]) add(key K, val V, new *bool) *node[K, V] {
	if n == nil {
		*new = true
		return newNode(key, val)
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
	} else if key < n.key {
		n.l = n.l.remove(key, removed, ok)
	} else {
		*removed = *n
		*ok = true
		defer func() { n = nil }()

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
		*n = rightMinNode
	}
	return n
}

func (n *node[K, V]) removeMax(removed *node[K, V], ok *bool) *node[K, V] {
	if n == nil {
		return nil
	}

	if n.r == nil {
		defer func() { n = nil }()
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
		defer func() { n = nil }()
		*ok = true
		*removed = *n
		return n.r
	}

	n.l = n.l.removeMin(removed, ok)
	return n
}

// XXX: Remove me
func (b *bTree[K, V]) debug(ts ...*testing.T) {
	q := []*node[K, V]{b.root}
	height := b.height() - 1
	descendOffset := descendOffset(height)
	descendTab := descendTab(height)
	println("")
	for i := 0; i < height; i++ {
		for j, l := 0, len(q); j < l; j++ {
			if j == 0 {
				print(strings.Repeat(" ", descendTab[i]))
			} else {
				print(strings.Repeat(" ", descendOffset[i]))
			}

			n := q[0]
			q = q[1:]
			if n != nil {
				print(n.val)
				q = append(q, n.l, n.r)
			} else {
				print("x")
				q = append(q, nil, nil)
			}
		}
		println("")
	}
	println("")
}

func (b *bTree[K, V]) height() int {
	var iter func(*node[K, V], int) int
	iter = func(n *node[K, V], count int) int {
		if n == nil {
			return count
		}
		return max(iter(n.l, count+1), iter(n.r, count+1))
	}
	return iter(b.root, 1)
}

func descendOffset(height int) []int {
	n := 0
	sli := make([]int, 0, height)
	for i := 0; i < height; i++ {
		n = n*2 + 1
		sli = append(sli, n)
	}

	result := make([]int, len(sli))
	for i := range sli {
		result[i] = sli[height-i-1]
	}
	return result
}

func descendTab(height int) []int {
	sli := make([]int, 0, height)
	for i := height - 1; i >= 0; i-- {
		sli = append(sli, pow(2, i))
	}
	return sli
}

func pow(n, x int) int {
	result := n
	if x <= 0 {
		return 1
	}

	for i := 1; i < x; i++ {
		result *= n
	}
	return result
}
