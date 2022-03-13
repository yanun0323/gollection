package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	assert.NotNil(t, s)
	assert.Equal(t, 0, s.Count())
	assert.True(t, s.IsEmpty())

	s2 := NewStack(data1, data2)
	assert.NotNil(t, s2)
	assert.Equal(t, 2, s2.Count())
	assert.False(t, s2.IsEmpty())
}
func Test_Stack_Clear(t *testing.T) {
	s := NewStack()
	s.Push(data1)
	s.Push(data2)
	s.Clear()

	assert.NotEqual(t, nil, s)
	assert.Equal(t, 0, s.Count())
}

func Test_Stack_Clone(t *testing.T) {
	s := NewStack()

	s.Push(data1)
	s.Push(data2)
	s.Push(data3)
	clone := s.Clone()
	es := s.Pop()
	ec := clone.Pop()
	assert.Equal(t, s.Count(), clone.Count())
	assert.Equal(t, es, ec)

	clone.Pop()
	assert.NotEqual(t, s.Count(), clone.Count())
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
	assert.Equal(t, 0, s.Count())

	s.Push(data1)
	assert.Equal(t, 1, s.Count())

	s.Pop()
	assert.Equal(t, 0, s.Count())
}

func Test_Stack_Pop(t *testing.T) {
	s := NewStack()
	assert.Panics(t, func() { s.Pop() })

	s.Push(data1)
	s.Push(data2)
	s.Push(data3)
	s.Push(nil)

	assert.Nil(t, s.Pop())
	assert.Equal(t, data3, s.Pop())
	assert.Equal(t, data2, s.Pop())
	assert.Equal(t, data1, s.Pop())

	assert.Panics(t, func() { s.Pop() })
}

func Test_Stack_Push(t *testing.T) {
	s := NewStack()

	s.Push(data1)
	assert.Equal(t, 1, s.Count())

	s.Push(data2)
	assert.Equal(t, 2, s.Count())

	s.Push(data3)
	assert.Equal(t, 3, s.Count())
}

func Test_Stack_IsEmpty(t *testing.T) {
	s := NewStack()

	assert.True(t, s.IsEmpty())

	s.Push(data1)
	assert.False(t, s.IsEmpty())
}

func Test_Stack_Peek(t *testing.T) {
	s := NewStack()
	assert.Panics(t, func() { s.Peek() })

	s.Push(data1)
	assert.Equal(t, data1, s.Peek())
	assert.Equal(t, 1, s.Count())
}

func Test_Stack_ToArray(t *testing.T) {
	s := NewStack()
	expect := []interface{}{data3, data2, data1}

	assert.Nil(t, s.ToArray())

	s.Push(data1)
	s.Push(data2)
	s.Push(data3)
	arr := s.ToArray()

	if assert.Equal(t, s.Count(), len(expect)) {
		for i := 0; i < s.Count(); i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}
}
