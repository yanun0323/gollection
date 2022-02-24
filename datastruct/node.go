package datastruct

type node struct {
	index    int
	data     *interface{}
	forward  *node
	backward *node
}

func newNode(i int, d interface{}, f *node, b *node) *node {
	if d == nil {
		return nil
	}
	return &node{index: i, data: &d, forward: f, backward: b}
}

type nodeF struct {
	data    *interface{}
	forward *nodeF
}

func newNodeF(d *interface{}, f *nodeF) *nodeF {
	if d == nil {
		return nil
	}
	return &nodeF{data: d, forward: f}
}

type nodeB struct {
	data     *interface{}
	backward *nodeB
}

func newNodeB(d *interface{}, b *nodeB) *nodeB {
	if d == nil {
		return nil
	}
	return &nodeB{data: d, backward: b}
}
