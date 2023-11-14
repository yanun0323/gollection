package gollection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBTree(t *testing.T) {
	suite.Run(t, new(BTreeSuite))
}

type BTreeSuite struct {
	suite.Suite
	mockTree func() *bTree[int, int]
}

/*
	    3
	2       5

1        4   6
*/
func (su *BTreeSuite) SetupTest() {
	su.mockTree = func() *bTree[int, int] {
		return &bTree[int, int]{
			count: 6,
			root: &node[int, int]{
				key: 3,
				val: 3,
				l: &node[int, int]{
					key: 2,
					val: 2,
					l: &node[int, int]{
						key: 1,
						val: 1,
					},
				},
				r: &node[int, int]{
					key: 5,
					val: 5,
					l: &node[int, int]{
						key: 4,
						val: 4,
					},
					r: &node[int, int]{
						key: 6,
						val: 6,
					},
				},
			},
		}
	}
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

func (su *BTreeSuite) Test_Len_Good() {
	b := su.mockTree()
	b.Insert(6, 6)
	su.Equal(7, b.Len())
	b.Insert(6, 6)
	su.Equal(8, b.Len())
	b.Insert(6, 6)
	su.Equal(9, b.Len())
	b.Insert(6, 6)
	su.Equal(10, b.Len())
}

func (su *BTreeSuite) Test_Insert_Good() {
	b := su.mockTree()
	b.Insert(3, 3)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.l)
	su.Require().NotNil(b.root.l.r)
	su.Equal(3, b.root.l.r.val)

	b.Insert(7, 7)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.r)
	su.Require().NotNil(b.root.r.r)
	su.Require().NotNil(b.root.r.r.r)
	su.Equal(7, b.root.r.r.r.val)

	e := bTree[int, int]{}
	e.Insert(5, 5)
	su.Require().NotNil(e.root)
	su.Equal(5, e.root.val)

	e.Insert(4, 4)
	su.Require().NotNil(e.root.l)
	su.Equal(4, e.root.l.val)

	e.Insert(3, 3)
	su.Require().NotNil(e.root.l.l)
	su.Equal(3, e.root.l.l.val)

	e.Insert(5, 5)
	su.Require().NotNil(e.root.l.r)
	su.Equal(5, e.root.l.r.val)

	e.Insert(6, 6)
	su.Require().NotNil(e.root.r)
	su.Equal(6, e.root.r.val)

	e.Insert(6, 6)
	su.Require().NotNil(e.root.r.l)
	su.Equal(6, e.root.r.l.val)

	e.Insert(7, 7)
	su.Require().NotNil(e.root.r.r)
	su.Equal(7, e.root.r.r.val)
}

func (su *BTreeSuite) Test_Remove_Good() {
	b := su.mockTree()
	r, ok := b.Remove(3)
	su.True(ok)
	su.Equal(3, r)
	su.check(b, 1, 2, 4, 5, 6)

	r, ok = b.Remove(5)
	su.True(ok)
	su.Equal(5, r)
	su.check(b, 1, 2, 4, 6)

	r, ok = b.Remove(4)
	su.True(ok)
	su.Equal(4, r)
	su.check(b, 1, 2, 6)

	r, ok = b.Remove(4)
	su.False(ok)
	su.Equal(0, r)
	su.check(b, 1, 2, 6)
}

func (su *BTreeSuite) check(b BTree[int, int], expected ...any) {
	i := 0
	su.Require().Equal(len(expected), b.Len())
	b.Ascend(func(k, v int) bool {
		su.Require().Equal(expected[i], k)
		i++
		return true
	})
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
