package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	assert.NotNil(t, s)
	assert.Equal(t, 0, s.count)
	assert.Nil(t, s.last)
}
func Test_Stack_Clear(t *testing.T) {
	s := NewStack()
	s.count = 2
	s.last = node2
	s.Clear()

	assert.NotEqual(t, nil, s)
	assert.Equal(t, 0, s.count)
	assert.Nil(t, s.last)
}

func Test_Stack_Clone(t *testing.T) {
	s := NewStack()

	s.Push(data1)
	s.Push(data2)
	clone := s.Clone()

	assert.Equal(t, s.Count(), clone.Count())
	assert.Equal(t, s.last, clone.last)

	clone.Pop()

	assert.NotEqual(t, s.count, clone.count)
}

func Test_Stack_Contains(t *testing.T) {
	s := NewStack()

	s.Push(data1)
	assert.True(t, s.Contains(data1))
	assert.True(t, s.Contains(data1, data3))
	assert.False(t, s.Contains(data2))
	assert.False(t, s.Contains(data3, data4))

	s.Push(data2)
	assert.True(t, s.Contains(data1))
	assert.True(t, s.Contains(data1, data2, data3))
	assert.True(t, s.Contains(data2))
	assert.False(t, s.Contains(data3, data4))
}

func Test_Stack_Count(t *testing.T) {
	s := NewStack()
	assert.Equal(t, s.count, s.Count())

	s.Push(data1)
	assert.Equal(t, s.count, s.Count())
}

func Test_Stack_Pop(t *testing.T) {
	s := NewStack()
	assert.Nil(t, s.Pop())

	s.Push(data1)
	s.Push(data2)
	s.Push(data3)

	assert.Equal(t, data3, s.Pop())
	assert.Equal(t, data2, s.Pop())
	assert.Equal(t, data1, s.Pop())
	assert.Nil(t, s.Pop())
}

func Test_Stack_Push(t *testing.T) {
	s := NewStack()
	s.Push(data1)
	assert.Equal(t, 1, s.count)
	assert.Equal(t, data1, *s.last.data)

	s.Push(data2)
	assert.Equal(t, 2, s.count)
	assert.Equal(t, data2, *s.last.data)

	s.Push(data3)
	assert.Equal(t, 3, s.count)
	assert.Equal(t, data3, *s.last.data)
}

func Test_Stack_IsEmpty(t *testing.T) {
	s := NewStack()

	assert.True(t, s.IsEmpty())

	s.Push(data1)
	assert.False(t, s.IsEmpty())
}

func Test_Stack_Peek(t *testing.T) {
	s := NewStack()

	assert.Nil(t, s.Peek())

	s.Push(data1)
	assert.NotNil(t, s.Peek())
}

func Test_Stack_ToSlice(t *testing.T) {
	s := NewStack()
	expect := []interface{}{data3, data2, data1}

	s.Push(data1)
	s.Push(data2)
	s.Push(data3)
	arr := s.ToArray()

	if assert.Equal(t, s.count, len(expect)) {
		for i := 0; i < s.count; i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}
}
