package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Count())
	assert.True(t, q.IsEmpty())
}
func Test_Queue_Clear(t *testing.T) {
	q := NewQueue()
	q.Enqueue(data1)
	q.Enqueue(data2)

	assert.True(t, q.Clear())

	assert.NotEqual(t, nil, q)
	assert.Equal(t, 0, q.Count())
}

func Test_Queue_Clone(t *testing.T) {
	q := NewQueue()

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	clone := q.Clone()
	eq, _ := q.Dequeue()
	ec, _ := clone.Dequeue()
	assert.Equal(t, q.Count(), clone.Count())
	assert.Equal(t, eq, ec)

	clone.Dequeue()
	assert.NotEqual(t, q.Count(), clone.Count())
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

	assert.False(t, q.Contains(nil))
	q.Enqueue(nil)
	assert.True(t, q.Contains(nil))
}

func Test_Queue_Count(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Count())

	q.Enqueue(data1)
	assert.Equal(t, 1, q.Count())

	q.Dequeue()
	assert.Equal(t, 0, q.Count())
}

func Test_Queue_Dequeue(t *testing.T) {
	q := NewQueue()
	_, ok := q.Dequeue()
	assert.False(t, ok)

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	q.Enqueue(nil)

	d, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, data1, d)

	d, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, data2, d)

	d, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, data3, d)

	d, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Nil(t, d)

	_, ok = q.Dequeue()
	assert.False(t, ok)
}

func Test_Queue_Enqueue(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.Enqueue(data1))
	assert.Equal(t, 1, q.Count())

	assert.True(t, q.Enqueue(data2))
	assert.Equal(t, 2, q.Count())

	assert.True(t, q.Enqueue(data3))
	assert.Equal(t, 3, q.Count())
}

func Test_Queue_IsEmpty(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.IsEmpty())

	q.Enqueue(data1)
	assert.False(t, q.IsEmpty())
}

func Test_Queue_Peek(t *testing.T) {
	q := NewQueue()
	_, ok := q.Peek()
	assert.False(t, ok)

	q.Enqueue(data1)
	d, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, data1, d)
	assert.Equal(t, 1, q.Count())
}

func Test_Queue_ToArray(t *testing.T) {
	q := NewQueue()
	expect := []interface{}{data1, data2, data3}

	_, ok := q.ToArray()
	assert.False(t, ok)

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	arr, ok := q.ToArray()

	assert.True(t, ok)
	if assert.Equal(t, q.Count(), len(expect)) {
		for i := 0; i < q.Count(); i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}
}
