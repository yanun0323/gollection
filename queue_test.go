package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data1 T = 10
	data2 T = 20
	data3 T = 30
	data4 T = 40

	node1 *node = newNode(&data1, nil, nil)
	node2 *node = newNode(&data2, nil, nil)
	node3 *node = newNode(&data3, nil, nil)
	node4 *node = newNode(&data4, nil, nil)
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	assert.NotNil(t, q)
	assert.Equal(t, 0, q.count)
	assert.Nil(t, q.first)
	assert.Nil(t, q.last)
}
func Test_Queue_Clear(t *testing.T) {
	q := NewQueue()
	q.count = 2
	q.first = node1
	q.last = node2
	q.Clear()

	assert.NotEqual(t, nil, q)
	assert.Equal(t, 0, q.count)
	assert.Nil(t, q.first)
	assert.Nil(t, q.last)
}

func Test_Queue_Clone(t *testing.T) {
	q := NewQueue()

	q.Enqueue(data1)
	q.Enqueue(data2)
	clone := q.Clone()

	assert.Equal(t, q.Count(), clone.Count())
	assert.Equal(t, q.first, clone.first)
	assert.Equal(t, q.last, clone.last)

	clone.Dequeue()

	assert.NotEqual(t, q.count, clone.count)
}

func Test_Queue_Contains(t *testing.T) {
	q := NewQueue()

	q.Enqueue(data1)
	assert.True(t, q.Contains(data1))
	assert.True(t, q.Contains(data1, data3))
	assert.False(t, q.Contains(data2))
	assert.False(t, q.Contains(data3, data4))

	q.Enqueue(data2)
	assert.True(t, q.Contains(data1))
	assert.True(t, q.Contains(data1, data2, data3))
	assert.True(t, q.Contains(data2))
	assert.False(t, q.Contains(data3, data4))
}

func Test_Queue_Count(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, q.count, q.Count())

	q.Enqueue(data1)
	assert.Equal(t, q.count, q.Count())
}

func Test_Queue_Dequeue(t *testing.T) {
	q := NewQueue()
	assert.Nil(t, q.Dequeue())

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)

	assert.Equal(t, data1, q.Dequeue())
	assert.Equal(t, data2, q.Dequeue())
	assert.Equal(t, data3, q.Dequeue())
	assert.Nil(t, q.Dequeue())
}

func Test_Queue_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(data1)
	assert.Equal(t, 1, q.count)
	assert.Equal(t, data1, *q.first.data)
	assert.Equal(t, data1, *q.last.data)

	q.Enqueue(data2)
	assert.Equal(t, 2, q.count)
	assert.Equal(t, data1, *q.first.data)
	assert.Equal(t, data2, *q.last.data)

	q.Enqueue(data3)
	assert.Equal(t, 3, q.count)
	assert.Equal(t, data1, *q.first.data)
	assert.Equal(t, data3, *q.last.data)
}

func Test_Queue_IsEmpty(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.IsEmpty())

	q.Enqueue(data1)
	assert.False(t, q.IsEmpty())
}

func Test_Queue_Peek(t *testing.T) {
	q := NewQueue()

	assert.Nil(t, q.Peek())

	q.Enqueue(data1)
	assert.NotNil(t, q.Peek())
}

func Test_Queue_ToArray(t *testing.T) {
	q := NewQueue()
	expect := []interface{}{data1, data2, data3}

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)

	if assert.Equal(t, q.count, len(expect)) {
		for i := 0; i < q.count; i++ {
			assert.Equal(t, expect[i], q.Dequeue())
		}
	}
}
