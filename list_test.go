package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_List_NewList(t *testing.T) {
	l := NewList()

	assert.NotNil(t, l)
	assert.Equal(t, 0, l.Count())
}

func Test_List_Add(t *testing.T) {
	l := NewList()

	l.ADD(data1)
	assert.NotNil(t, l)
	assert.Equal(t, 1, l.Count())

	l.ADD(data2, data3, nil)
	assert.NotNil(t, l)
	assert.Equal(t, 4, l.Count())
	d, _ := l.At(0)
	assert.Equal(t, data1, d)
	d, _ = l.At(1)
	assert.Equal(t, data2, d)
	d, _ = l.At(2)
	assert.Equal(t, data3, d)
	d, _ = l.At(3)
	assert.Nil(t, d)

}

func Test_List_At(t *testing.T) {
	l := NewList()

	l.ADD(data1, data2, data3)
	d, _ := l.At(0)
	assert.Equal(t, data1, d)
	d, _ = l.At(1)
	assert.Equal(t, data2, d)
	d, _ = l.At(2)
	assert.Equal(t, data3, d)
	d, _ = l.At(3)
	assert.Nil(t, d)
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
	d, _ := l.At(0)
	dc, _ := clone.At(0)
	assert.Equal(t, dc, d)
	d, _ = l.At(1)
	dc, _ = clone.At(1)
	assert.Equal(t, dc, d)
	d, _ = l.At(2)
	dc, _ = clone.At(2)
	assert.Equal(t, dc, d)
	d, _ = l.At(3)
	dc, _ = clone.At(3)
	assert.Equal(t, dc, d)
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
	d, _ := l.At(0)
	assert.Equal(t, data1, d)

	assert.True(t, l.Insert(2, data3))
	d, _ = l.At(2)
	assert.Equal(t, data3, d)

	assert.True(t, l.Insert(4, data5))
	d, _ = l.At(4)
	assert.Equal(t, data5, d)

	assert.False(t, l.Insert(-1, data5))
	assert.False(t, l.Insert(6, data5))
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
	d, _ := l.At(0)
	assert.Equal(t, data2, d)
	assert.Equal(t, 4, l.Count())
}

func Test_List_RemoveAll(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data1, data3, data1)

	assert.False(t, l.RemoveAll(data4))
	assert.True(t, l.RemoveAll(data1))
	d, _ := l.At(0)
	assert.Equal(t, data2, d)
	assert.Equal(t, 2, l.Count())
}

func Test_List_RemoveAt(t *testing.T) {
	l := NewList()
	l.ADD(data1, data2, data1, data3, data1)

	d, ok := l.RemoveAt(4)
	assert.True(t, ok)
	assert.Equal(t, data1, d)
	assert.Equal(t, 4, l.Count())

	d, ok = l.RemoveAt(2)
	assert.True(t, ok)
	assert.Equal(t, data1, d)
	assert.Equal(t, 3, l.Count())

}

func Test_List_ToArray(t *testing.T) {
	l := NewList()
	expect := []interface{}{data1, data2, data1, data3, data1}

	_, ok := l.ToArray()
	assert.False(t, ok)

	l.ADD(data1, data2, data1, data3, data1)

	arr, ok := l.ToArray()
	assert.True(t, ok)
	if assert.Equal(t, l.Count(), len(expect)) {
		for i := 0; i < l.Count(); i++ {
			assert.Equal(t, expect[i], arr[i])
		}
	}

}
