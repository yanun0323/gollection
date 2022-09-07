package gollection

type bTree struct {
	count   int
	root    *node
	greater func(any, any) bool
}

func NewBTree(greater func(any, any) bool) bTree {
	return bTree{
		greater: greater,
	}
}

func (b *bTree) Count() int {
	return b.count
}

func (b *bTree) Insert(v any) {
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
func (b *bTree) Walk(t int) []any {
	result := make([]any, 0, b.count)
	f := b.getOrder(t)
	f(&result, b.root)
	return result
}

func (b *bTree) getOrder(t int) func(*[]any, *node) {
	if t < 0 {
		return b.preOrder
	}
	if t == 0 {
		return b.inOrder
	}
	return b.postOrder
}

// Root -> L -> R
func (b *bTree) preOrder(sli *[]any, n *node) {
	if n == nil {
		return
	}
	*sli = append(*sli, n.val)
	b.preOrder(sli, n.l)
	b.preOrder(sli, n.r)
}

// L -> Root -> R
func (b *bTree) inOrder(sli *[]any, n *node) {
	if n == nil {
		return
	}
	b.inOrder(sli, n.l)
	*sli = append(*sli, n.val)
	b.inOrder(sli, n.r)
}

// L -> R -> Root
func (b *bTree) postOrder(sli *[]any, n *node) {
	if n == nil {
		return
	}
	b.postOrder(sli, n.l)
	b.postOrder(sli, n.r)
	*sli = append(*sli, n.val)
}

type node struct {
	val any
	l   *node
	r   *node
}

func newNode(a any) *node {
	return &node{
		val: a,
	}
}

func (n *node) Insert(nn *node, greater func(any, any) bool) {
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
