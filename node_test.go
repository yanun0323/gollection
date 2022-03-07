package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data1 interface{} = 10
	data2 interface{} = 20
	data3 interface{} = 30
	data4 interface{} = 40

	node1 *node = newNode(&data1, nil, nil)
	node2 *node = newNode(&data2, nil, nil)
)

func Test_Node_NewNode(t *testing.T) {
	var data interface{} = 20
	node_data := newNode(&data, nil, nil)
	node_nil := newNode(nil, nil, nil)

	assert.NotNil(t, node_data)
	assert.Nil(t, node_nil)
}
