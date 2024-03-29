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
			m: map[int]struct{}{
				0: {},
				1: {},
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
				9: {},
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
	su.NotNil(s.m[10])
	su.Equal(11, len(s.m))

	s.Insert(11)
	s.Insert(12)
	s.Insert(13)
	su.NotNil(s.m[11])
	su.NotNil(s.m[12])
	su.NotNil(s.m[13])
	su.Equal(14, len(s.m))

	s.Insert(0)
	su.Equal(14, len(s.m))
}

func (su *SetSuite) Test_Remove_Good() {
	s := su.mockSet()
	s.Remove(10)
	su.Equal(10, len(s.m))

	s.Remove(9)
	su.Equal(9, len(s.m))

	s.Remove(6, 7, 8)
	su.Equal(6, len(s.m))

	e := set[int]{}
	e.Remove(0)
	su.Equal(0, len(e.m))
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
	su.Equal(len(s.m), s.Len())

	e := NewSet[int]()
	su.Equal(0, e.Len())
}
