package gollection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestStack(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

type StackSuite struct {
	suite.Suite
	mockStack func() stack
}

func (su *StackSuite) SetupTest() {
	su.mockStack = func() stack {
		return stack{
			data: []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
	}
}

func (su *StackSuite) Test_NewStack_Good() {
	s := NewStack()
	su.Equal(0, len(s.data))
}

func (su *StackSuite) Test_Push_Good() {
	s := su.mockStack()
	s.Push(10)
	su.Equal(10, s.data[10])
	s.Push(11, 12, 13)
	su.Equal(11, s.data[11])
	su.Equal(12, s.data[12])
	su.Equal(13, s.data[13])
}

func (su *StackSuite) Test_Pop_Good() {
	s := su.mockStack()
	su.Equal(9, s.Pop())
	su.Equal(8, s.data[len(s.data)-1])
	su.Equal(9, len(s.data))

	e := stack{}
	su.Nil(e.Pop())
}

func (su *StackSuite) Test_Peek_Good() {
	s := su.mockStack()
	su.Equal(9, s.Peek())
	su.Equal(9, s.data[len(s.data)-1])
	su.Equal(10, len(s.data))

	e := stack{}
	su.Nil(e.Peek())
}

func (su *StackSuite) Test_Count_Good() {
	s := su.mockStack()
	su.Equal(len(s.data), s.Len())

	e := stack{}
	su.Equal(0, e.Len())
}

func (su *StackSuite) Test_ToSlice_Good() {
	s := su.mockStack()
	sli := s.ToSlice()
	su.Require().Equal(len(s.data), len(sli))
	for i := range sli {
		su.Equal(s.data[len(s.data)-1-i], sli[i])
	}

	sli[0] = 123
	su.Equal(0, s.data[0])
	su.Equal(9, s.data[9])
	su.Equal(123, sli[0])
}
