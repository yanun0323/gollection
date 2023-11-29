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

func (su *BTreeSuite) check(b BTree[int, int], expected ...any) {
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
	su.Equal(2, b.root.l.r.val)

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
	b.debug()

	println("remove 1")
	r, ok := b.Remove(1)
	su.Require().True(ok)
	su.Require().Equal(1, r)
	su.check(b, 2, 3, 4, 5, 6)

	b.debug()

	println("remove 3")
	r, ok = b.Remove(3)
	su.Require().True(ok)
	su.Require().Equal(3, r)
	su.check(b, 2, 4, 5, 6)

	b.debug()

	println("remove 5")
	r, ok = b.Remove(5)
	su.Require().True(ok)
	su.Require().Equal(5, r)
	su.check(b, 2, 4, 6)

	b.debug()

	println("remove 4")
	r, ok = b.Remove(4)
	su.Require().True(ok)
	su.Require().Equal(4, r)
	su.check(b, 2, 6)

	b.debug()

	println("remove 4")
	r, ok = b.Remove(4)
	su.Require().False(ok)
	su.Require().Equal(0, r)
	su.check(b, 2, 6)

	b.debug()
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
	su.Equal(6, b.Len())
	b.debug()

	k, v, ok = b.RemoveMin()
	su.True(ok)
	su.Equal(1, k)
	su.Equal(1, v)
	su.Equal(5, b.Len())
	println("remove min")
	b.debug()

	v, ok = b.Remove(1)
	su.False(ok)
	su.Equal(0, v)
	println("remove 1")
	b.debug()

	k, v, ok = b.RemoveMin()
	su.True(ok)
	su.Equal(2, k)
	su.Equal(2, v)
	su.Equal(4, b.Len())
	println("remove min")
	b.debug()
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
