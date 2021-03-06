package gollection

type node struct {
	data     *interface{}
	forward  *node
	backward *node
}

func newNode(data *interface{}, forward *node, backward *node) *node {
	if data == nil {
		return nil
	}
	return &node{data: data, forward: forward, backward: backward}
}
