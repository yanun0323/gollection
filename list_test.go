package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_List_NewList(t *testing.T) {
	l := NewList()

	assert.NotNil(t, l)
	assert.Equal(t, 0, l.Count())

	l2 := NewList(data1, data2)

	assert.NotNil(t, l2)
	assert.Equal(t, 2, l2.Count())
	assert.Equal(t, data1, l2.At(0))
	assert.Equal(t, data2, l2.At(1))
}

func Test_List_Add(t *testing.T) {
	l := NewList()

	l.ADD(data1)
	assert.NotNil(t, l)
	assert.Equal(t, 1, l.Count())

	l.ADD(data2, data3, nil)
	assert.NotNil(t, l)
	assert.Equal(t, 4, l.Count())
	assert.Equal(t, data1, l.At(0))
	assert.Equal(t, data2, l.At(1))
	assert.Equal(t, data3, l.At(2))
	assert.Nil(t, l.At(3))

}

func Test_List_At(t *testing.T) {
	l := NewList()

	l.ADD(data1, data2, data3)

	assert.Equal(t, data1, l.At(0))
	assert.Equal(t, data2, l.At(1))
	assert.Equal(t, data3, l.At(2))
	assert.Panics(t, func() { l.At(3) })
	assert.Panics(t, func() { l.At(-1) })
}

func Test_List_Clear(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data3)
	l.Clear()

	assert.NotNil(t, l)
	assert.Equal(t, 0, l.Count())
}
func Test_List_Clone(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data3)
	clone := l.Clone()

	assert.NotNil(t, clone)
	assert.Equal(t, l.Count(), clone.Count())
	assert.Equal(t, clone.At(0), l.At(0))
	assert.Equal(t, clone.At(1), l.At(1))
	assert.Equal(t, clone.At(2), l.At(2))
}

func Test_List_Contains(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data3)

	assert.True(t, l.Contains(data1))
	assert.True(t, l.Contains(data2, data3))
	assert.True(t, l.Contains(data2, data3, data4))
	assert.False(t, l.Contains(data4))
	assert.False(t, l.Contains(data4, data5))
}

func Test_List_Count(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data3)

	assert.True(t, l.Contains(data1))
	assert.True(t, l.Contains(data2, data3))
	assert.True(t, l.Contains(data2, data3, data4))
	assert.False(t, l.Contains(data4))
	assert.False(t, l.Contains(data4, data5))
}

func Test_List_Insert(t *testing.T) {
	l := NewList()
	l.ADD(data2, data4)

	assert.True(t, l.Insert(0, data1))
	assert.Equal(t, data1, l.At(0))

	assert.True(t, l.Insert(2, data3))
	assert.Equal(t, data3, l.At(2))

	assert.True(t, l.Insert(4, data5))
	assert.Equal(t, data5, l.At(4))

	assert.Panics(t, func() { l.Insert(-1, data5) })
	assert.Panics(t, func() { l.Insert(6, data5) })
}

func Test_List_IsEmpty(t *testing.T) {
	l := NewList()

	assert.True(t, l.IsEmpty())

	l.ADD(data2, data4)
	assert.False(t, l.IsEmpty())
}

func Test_List_Remove(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data1, data3, data1)

	assert.False(t, l.Remove(data4))
	assert.True(t, l.Remove(data1))

	assert.Equal(t, data2, l.At(0))
	assert.Equal(t, 4, l.Count())
}

func Test_List_RemoveAll(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data1, data3, data1)

	assert.False(t, l.RemoveAll(data4))
	assert.True(t, l.RemoveAll(data1))

	assert.Equal(t, data2, l.At(0))
	assert.Equal(t, 2, l.Count())
}

func Test_List_RemoveAt(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data1, data3, data1)

	assert.Equal(t, data1, l.RemoveAt(4))
	assert.Equal(t, 4, l.Count())

	assert.Equal(t, data1, l.RemoveAt(2))
	assert.Equal(t, 3, l.Count())

	assert.Panics(t, func() { l.RemoveAt(-1) })
	assert.Panics(t, func() { l.RemoveAt(6) })

}

func Test_List_ToArray(t *testing.T) {
	l := NewList()
	expect := []interface{}{data1, data2, data1, data3, data1}

	assert.Nil(t, l.ToArray())

	l.ADD(data1, data2, data1, data3, data1)

	arr := l.ToArray()
	if assert.Equal(t, l.Count(), len(expect)) {
		for i := 0; i < l.Count(); i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}

}
