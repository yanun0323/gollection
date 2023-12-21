package gollection

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBTree(t *testing.T) {
	suite.Run(t, new(BTreeSuite))
}

type BTreeSuite struct {
	suite.Suite
	// default content: 1, 2, 3, 4, 5, 6
	mockTree func() *bTree[int, int]
}

/*
	    3
	2       5

1        4   6
*/
func (su *BTreeSuite) SetupTest() {
	su.mockTree = func() *bTree[int, int] {
		t := &bTree[int, int]{}
		t.Insert(3, 3)
		t.Insert(1, 1)
		t.Insert(2, 2)
		t.Insert(5, 5)
		t.Insert(4, 4)
		t.Insert(6, 6)
		return t
	}
}

func (su *BTreeSuite) check(b *bTree[int, int], expected ...any) {
	i := 0
	su.Require().Equal(len(expected), b.Len())
	b.Ascend(func(k, v int) bool {
		su.Require().Equal(expected[i], k)
		i++
		return true
	})
}

func (su *BTreeSuite) Test_NewBTree_Good() {
	b := NewBTree[int, int]()
	su.NotNil(b)

	su.Require().NotPanics(func() {
		b.Insert(1, 1)
		b.Insert(2, 2)
	})
	su.Equal(2, b.Len())
}

func (su *BTreeSuite) Test_Contain_Good() {
	b := su.mockTree()
	su.True(b.Contain(1))
	su.True(b.Contain(5))
	su.False(b.Contain(9))
}

func (su *BTreeSuite) Test_Len_Good() {
	b := su.mockTree()
	b.Insert(7, 6)
	su.Equal(7, b.Len())
	b.Insert(8, 6)
	su.Equal(8, b.Len())
	b.Insert(9, 6)
	su.Equal(9, b.Len())
	b.Insert(9, 6)
	su.Equal(9, b.Len())
}

func (su *BTreeSuite) Test_Insert_Good() {
	b := su.mockTree()
	b.Insert(3, 3)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.l)
	su.Require().NotNil(b.root.l.r)
	su.Equal(2, b.root.l.r.val)

	b.Insert(7, 7)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.r)
	su.Require().NotNil(b.root.r.r)
	su.Require().NotNil(b.root.r.r.r)
	su.Equal(7, b.root.r.r.r.val)

	e := bTree[int, int]{}
	e.Insert(7, 7)
	su.Require().NotNil(e.root)
	su.Equal(7, e.root.val)

	e.Insert(4, 4)
	su.Require().NotNil(e.root.l)
	su.Equal(4, e.root.l.val)

	e.Insert(3, 3)
	su.Require().NotNil(e.root.l.l)
	su.Equal(3, e.root.l.l.val)

	e.Insert(5, 5)
	su.Require().NotNil(e.root.l.r)
	su.Equal(5, e.root.l.r.val)

	e.Insert(10, 10)
	su.Require().NotNil(e.root.r)
	su.Equal(10, e.root.r.val)

	e.Insert(8, 8)
	su.Require().NotNil(e.root.r.l)
	su.Equal(8, e.root.r.l.val)

	e.Insert(11, 11)
	su.Require().NotNil(e.root.r.r)
	su.Equal(11, e.root.r.r.val)
}

func (su *BTreeSuite) Test_Remove_Good() {
	b := su.mockTree()
	su.debug(b, "init")

	r, ok := b.Remove(1)
	su.debug(b, "remove 1")
	su.Require().True(ok)
	su.Require().Equal(1, r)
	su.check(b, 2, 3, 4, 5, 6)

	r, ok = b.Remove(3)
	su.debug(b, "remove 3")
	su.Require().True(ok)
	su.Require().Equal(3, r)
	su.check(b, 2, 4, 5, 6)

	r, ok = b.Remove(5)
	su.debug(b, "remove 5")
	su.Require().True(ok)
	su.Require().Equal(5, r)
	su.check(b, 2, 4, 6)

	r, ok = b.Remove(4)
	su.debug(b, "remove 4")
	su.Require().True(ok)
	su.Require().Equal(4, r)
	su.check(b, 2, 6)

	r, ok = b.Remove(4)
	su.debug(b, "remove 4")
	su.Require().False(ok)
	su.Require().Equal(0, r)
	su.check(b, 2, 6)
}

func (su *BTreeSuite) Test_Search_Good() {
	b := su.mockTree()
	testCases := []struct {
		desc    string
		key     int
		success bool
	}{
		{
			key:     9,
			success: false,
		},
		{
			key:     -1,
			success: false,
		},
		{
			key:     3,
			success: true,
		},
	}

	for _, tc := range testCases {
		su.T().Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			v, ok := b.Search(tc.key)
			if tc.success {
				su.True(ok)
				su.Equal(tc.key, v)
			} else {
				su.False(ok)
				su.Zero(v)
			}
		})
	}
}

func (su *BTreeSuite) Test_Max_Good() {
	bt := NewBTree[int, int]()
	k, v, ok := bt.Max()
	su.False(ok)
	su.Zero(k)
	su.Zero(v)

	b := su.mockTree()
	k, v, ok = b.Max()
	su.True(ok)
	su.Equal(6, k)
	su.Equal(6, v)

	v, ok = b.Remove(6)
	su.True(ok)
	su.Equal(6, v)

	k, v, ok = b.Max()
	su.True(ok)
	su.Equal(5, k)
	su.Equal(5, v)
}

func (su *BTreeSuite) Test_Min_Good() {
	bt := NewBTree[int, int]()
	k, v, ok := bt.Min()
	su.False(ok)
	su.Zero(k)
	su.Zero(v)

	b := su.mockTree()
	k, v, ok = b.Min()
	su.True(ok)
	su.Equal(1, k)
	su.Equal(1, v)

	v, ok = b.Remove(1)
	su.True(ok)
	su.Equal(1, v)

	k, v, ok = b.Min()
	su.True(ok)
	su.Equal(2, k)
	su.Equal(2, v)
}

func (su *BTreeSuite) Test_RemoveMax_Good() {
	bt := NewBTree[int, int]()
	su.Equal(0, bt.Len())

	k, v, ok := bt.RemoveMax()
	su.False(ok)
	su.Zero(k)
	su.Zero(v)
	su.Equal(0, bt.Len())

	b := su.mockTree()
	su.Equal(6, b.Len())

	k, v, ok = b.RemoveMax()
	su.True(ok)
	su.Equal(6, k)
	su.Equal(6, v)
	su.Equal(5, b.Len())

	v, ok = b.Remove(6)
	su.False(ok)
	su.Zero(v)
	su.Equal(5, b.Len())

	k, v, ok = b.RemoveMax()
	su.True(ok)
	su.Equal(5, k)
	su.Equal(5, v)
	su.Equal(4, b.Len())
}

func (su *BTreeSuite) Test_RemoveMin_Good() {
	bt := &bTree[int, int]{}
	su.Equal(0, bt.Len())

	k, v, ok := bt.RemoveMin()
	su.False(ok)
	su.Zero(k)
	su.Zero(v)
	su.Equal(0, bt.Len())

	b := su.mockTree()
	su.debug(b, "init")
	su.Equal(6, b.Len())

	k, v, ok = b.RemoveMin()
	su.debug(b, "remove min")
	su.True(ok)
	su.Equal(1, k)
	su.Equal(1, v)
	su.Equal(5, b.Len())

	v, ok = b.Remove(1)
	su.debug(b, "remove 1")
	su.False(ok)
	su.Equal(0, v)

	k, v, ok = b.RemoveMin()
	su.debug(b, "remove min")
	su.True(ok)
	su.Equal(2, k)
	su.Equal(2, v)
	su.Equal(4, b.Len())
}

func (su *BTreeSuite) Test_Ascend_Good() {
	b := su.mockTree()
	expected := []any{1, 2, 3, 4, 5, 6}
	i := 0
	su.Require().Equal(len(expected), b.Len())
	b.Ascend(func(k, v int) bool {
		su.Require().Equal(expected[i], k)
		i++
		return true
	})
}

func (su *BTreeSuite) Test_Descend_Good() {
	b := su.mockTree()
	expected := []any{6, 5, 4, 3, 2, 1}
	i := 0
	su.Require().Equal(len(expected), b.Len())
	b.Descend(func(k, v int) bool {
		su.Require().Equal(expected[i], k)
		i++
		return true
	})
}

func (su *BTreeSuite) debug(b *bTree[int, int], name string) {
	su.T().Log(name)
	q := []*node[int, int]{b.root}
	height := height(b) - 1
	descendOffset := descendOffset(height)
	descendTab := descendTab(height)
	buf := strings.Builder{}
	buf.WriteByte('\n')
	for i := 0; i < height; i++ {
		for j, l := 0, len(q); j < l; j++ {
			if j == 0 {
				buf.WriteString(strings.Repeat(" ", descendTab[i]))
			} else {
				buf.WriteString(strings.Repeat(" ", descendOffset[i]))
			}

			n := q[0]
			q = q[1:]
			if n != nil {
				buf.WriteString(fmt.Sprintf("%d", n.val))
				q = append(q, n.l, n.r)
			} else {
				buf.WriteString("x")
				q = append(q, nil, nil)
			}
		}
		buf.WriteByte('\n')
	}
	su.T().Log(buf.String())
}

func height(b *bTree[int, int]) int {
	var iter func(*node[int, int], int) int
	iter = func(n *node[int, int], count int) int {
		if n == nil {
			return count
		}
		return max(iter(n.l, count+1), iter(n.r, count+1))
	}
	return iter(b.root, 1)
}

func descendOffset(height int) []int {
	n := 0
	sli := make([]int, 0, height)
	for i := 0; i < height; i++ {
		n = n*2 + 1
		sli = append(sli, n)
	}

	result := make([]int, len(sli))
	for i := range sli {
		result[i] = sli[height-i-1]
	}
	return result
}

func descendTab(height int) []int {
	sli := make([]int, 0, height)
	for i := height - 1; i >= 0; i-- {
		sli = append(sli, pow(2, i))
	}
	return sli
}

func pow(n, x int) int {
	result := n
	if x <= 0 {
		return 1
	}

	for i := 1; i < x; i++ {
		result *= n
	}
	return result
}
