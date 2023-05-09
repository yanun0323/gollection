package gollection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSet(t *testing.T) {
	suite.Run(t, new(SetSuite))
}

type SetSuite struct {
	suite.Suite
	mockSet func() set[int]
}

func (su *SetSuite) SetupTest() {
	su.mockSet = func() set[int] {
		return set[int]{
			hash: map[int]bool{
				0: true,
				1: true,
				2: true,
				3: true,
				4: true,
				5: true,
				6: true,
				7: true,
				8: true,
				9: true,
			},
		}
	}
}

func (su *SetSuite) Test_NewSet_Good() {
	s := NewSet[int]()
	su.NotNil(s)
}

func (su *SetSuite) Test_Insert_Good() {
	s := su.mockSet()
	s.Insert(10)
	su.True(s.hash[10])
	su.Equal(11, len(s.hash))

	s.Insert(11)
	s.Insert(12)
	s.Insert(13)
	su.True(s.hash[11])
	su.True(s.hash[12])
	su.True(s.hash[13])
	su.Equal(14, len(s.hash))

	s.Insert(0)
	su.Equal(14, len(s.hash))
}

func (su *SetSuite) Test_Remove_Good() {
	s := su.mockSet()
	s.Remove(10)
	su.Equal(10, len(s.hash))

	s.Remove(9)
	su.Equal(9, len(s.hash))

	s.Remove(6, 7, 8)
	su.Equal(6, len(s.hash))

	e := set[int]{}
	e.Remove(0)
	su.Equal(0, len(e.hash))
}

func (su *SetSuite) Test_Contain_Good() {
	s := su.mockSet()
	su.True(s.Contain(0))
	su.True(s.Contain(5))
	su.True(s.Contain(9))
	su.False(s.Contain(10))

	e := NewSet[int]()
	su.False(e.Contain(0))
}

func (su *SetSuite) Test_Count_Good() {
	s := su.mockSet()
	su.Equal(len(s.hash), s.Len())

	e := NewSet[int]()
	su.Equal(0, e.Len())
}
