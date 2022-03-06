package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node_NewNode(t *testing.T) {
	var data T = 20
	node_data := newNode(&data, nil, nil)
	node_nil := newNode(nil, nil, nil)

	assert.NotNil(t, node_data)
	assert.Nil(t, node_nil)
}
