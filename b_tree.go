package gollection

type BTree[T any] interface {
	Len() int
	Insert(v T)
	/*
		walk through the tree

			t < 0  : Pre-Order Traversal
			t == 0 : In-Order Traversal
			t > 0  : Post-Order Traversal
	*/
	Walk(t int) []T
}

type bTree[T any] struct {
	count   int
	root    *node[T]
	greater func(T, T) bool
}

func NewBTree[T any](greater func(t1, t2 T) bool) BTree[T] {
	return &bTree[T]{
		greater: greater,
	}
}

func (b *bTree[T]) Len() int {
	return b.count
}

func (b *bTree[T]) Insert(v T) {
	b.count++
	if b.root == nil {
		b.root = newNode(v)
		return
	}
	b.root.Insert(newNode(v), b.greater)
}

/*
walk through the tree

	t < 0  : Pre-Order Traversal
	t == 0 : In-Order Traversal
	t > 0  : Post-Order Traversal
*/
func (b *bTree[T]) Walk(t int) []T {
	result := make([]T, 0, b.count)
	f := b.getOrder(t)
	f(&result, b.root)
	return result
}

func (b *bTree[T]) getOrder(t int) func(*[]T, *node[T]) {
	if t < 0 {
		return b.preOrder
	}
	if t == 0 {
		return b.inOrder
	}
	return b.postOrder
}

// Root -> L -> R
func (b *bTree[T]) preOrder(sli *[]T, n *node[T]) {
	if n == nil {
		return
	}
	*sli = append(*sli, n.val)
	b.preOrder(sli, n.l)
	b.preOrder(sli, n.r)
}

// L -> Root -> R
func (b *bTree[T]) inOrder(sli *[]T, n *node[T]) {
	if n == nil {
		return
	}
	b.inOrder(sli, n.l)
	*sli = append(*sli, n.val)
	b.inOrder(sli, n.r)
}

// L -> R -> Root
func (b *bTree[T]) postOrder(sli *[]T, n *node[T]) {
	if n == nil {
		return
	}
	b.postOrder(sli, n.l)
	b.postOrder(sli, n.r)
	*sli = append(*sli, n.val)
}

type node[T any] struct {
	val T
	l   *node[T]
	r   *node[T]
}

func newNode[T any](a T) *node[T] {
	return &node[T]{
		val: a,
	}
}

func (n *node[T]) Insert(nn *node[T], greater func(T, T) bool) {
	switch greater(nn.val, n.val) {
	case true:
		if n.r == nil {
			n.r = nn
			return
		}
		n.r.Insert(nn, greater)
	case false:
		if n.l == nil {
			n.l = nn
			return
		}
		n.l.Insert(nn, greater)
	}
}
