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

	q2 := NewQueue(data1, data2)

	assert.NotNil(t, q2)
	assert.Equal(t, 2, q2.Count())
	assert.False(t, q2.IsEmpty())
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
	eq := q.Dequeue()
	ec := clone.Dequeue()
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
	assert.Panics(t, func() { q.Dequeue() })

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	q.Enqueue(nil)

	assert.Equal(t, data1, q.Dequeue())
	assert.Equal(t, data2, q.Dequeue())
	assert.Equal(t, data3, q.Dequeue())
	assert.Nil(t, q.Dequeue())

	assert.Panics(t, func() { q.Dequeue() })
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

	assert.Panics(t, func() {
		q.Peek()
	})

	q.Enqueue(data1)
	assert.Equal(t, data1, q.Peek())
	assert.Equal(t, 1, q.Count())
}

func Test_Queue_ToArray(t *testing.T) {
	q := NewQueue()
	expect := []interface{}{data1, data2, data3}

	assert.Nil(t, q.ToArray())

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	arr := q.ToArray()

	if assert.Equal(t, q.Count(), len(expect)) {
		for i := 0; i < q.Count(); i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}
}
