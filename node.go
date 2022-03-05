package gollection

type node struct {
	data     *T
	forward  *node
	backward *node
}

func newNode(data T, forward *node, backward *node) *node {
	if data == nil {
		return nil
	}
	return &node{data: &data, forward: forward, backward: backward}
}
