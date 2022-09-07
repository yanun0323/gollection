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
	mockTree func() bTree
}

/*
       3
   2       5
1        4   6
*/
func (su *BTreeSuite) SetupTest() {
	su.mockTree = func() bTree {
		return bTree{
			count: 6,
			root: &node{
				val: 3,
				l: &node{
					val: 2,
					l: &node{
						val: 1,
					},
				},
				r: &node{
					val: 5,
					l: &node{
						val: 4,
					},
					r: &node{
						val: 6,
					},
				},
			},
			greater: func(a1, a2 any) bool {
				return a1.(int) > a2.(int)
			},
		}
	}
}

func (su *BTreeSuite) Test_NewBTree_Good() {
	b := NewBTree(func(a1, a2 any) bool {
		return a1.(int) > a2.(int)
	})
	su.Equal(0, b.count)
	su.Nil(b.root)
	su.True(b.greater(2, 1))
	su.False(b.greater(1, 1))
	su.False(b.greater(1, 2))
}

func (su *BTreeSuite) Test_Count_Good() {
	b := su.mockTree()
	b.Insert(6)
	su.Equal(7, b.Count())
	b.Insert(6)
	su.Equal(8, b.Count())
	b.Insert(6)
	su.Equal(9, b.Count())
	b.Insert(6)
	su.Equal(10, b.Count())
}

func (su *BTreeSuite) Test_Insert_Good() {
	b := su.mockTree()
	b.Insert(3)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.l)
	su.Require().NotNil(b.root.l.r)
	su.Equal(3, b.root.l.r.val)

	b.Insert(7)
	su.Require().NotNil(b.root)
	su.Require().NotNil(b.root.r)
	su.Require().NotNil(b.root.r.r)
	su.Require().NotNil(b.root.r.r.r)
	su.Equal(7, b.root.r.r.r.val)

	e := bTree{
		greater: func(a1, a2 any) bool {
			return a1.(int) > a2.(int)
		},
	}
	e.Insert(5)
	su.Require().NotNil(e.root)
	su.Equal(5, e.root.val)

	e.Insert(4)
	su.Require().NotNil(e.root.l)
	su.Equal(4, e.root.l.val)

	e.Insert(3)
	su.Require().NotNil(e.root.l.l)
	su.Equal(3, e.root.l.l.val)

	e.Insert(5)
	su.Require().NotNil(e.root.l.r)
	su.Equal(5, e.root.l.r.val)

	e.Insert(6)
	su.Require().NotNil(e.root.r)
	su.Equal(6, e.root.r.val)

	e.Insert(6)
	su.Require().NotNil(e.root.r.l)
	su.Equal(6, e.root.r.l.val)

	e.Insert(7)
	su.Require().NotNil(e.root.r.r)
	su.Equal(7, e.root.r.r.val)
}

func (su *BTreeSuite) Test_Walk_PreOrder_Good() {
	b := su.mockTree()
	result := b.Walk(-1)
	expected := []any{3, 2, 1, 5, 4, 6}
	su.Require().Equal(len(expected), len(result))
	for i := range result {
		su.Equal(expected[i], result[i])
	}
}

func (su *BTreeSuite) Test_Walk_InOrder_Good() {
	b := su.mockTree()
	result := b.Walk(0)
	expected := []any{1, 2, 3, 4, 5, 6}
	su.Require().Equal(len(expected), len(result))
	for i := range result {
		su.Equal(expected[i], result[i])
	}

}

func (su *BTreeSuite) Test_Walk_PostOrder_Good() {
	b := su.mockTree()
	result := b.Walk(1)
	expected := []any{1, 2, 4, 6, 5, 3}
	su.Require().Equal(len(expected), len(result))
	for i := range result {
		su.Equal(expected[i], result[i])
	}
}
