package gollection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestQueue(t *testing.T) {
	suite.Run(t, new(QueueSuite))
}

type QueueSuite struct {
	suite.Suite
	mockQueue func() queue[int]
}

func (su *QueueSuite) SetupTest() {
	su.mockQueue = func() queue[int] {
		return queue[int]{
			data: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
	}
}

func (su *QueueSuite) Test_NewQueue_Good() {
	q := NewQueue[int]()
	su.NotNil(q)
}

func (su *QueueSuite) Test_Enqueue_Good() {
	q := su.mockQueue()
	q.Enqueue(10)
	su.Equal(10, q.data[10])
	q.Enqueue(11, 12, 13)
	su.Equal(11, q.data[11])
	su.Equal(12, q.data[12])
	su.Equal(13, q.data[13])
}

func (su *QueueSuite) Test_Dequeue_Good() {
	q := su.mockQueue()
	su.Equal(-1, q.Dequeue())
	su.Equal(1, q.data[0])
	su.Equal(9, len(q.data))

	e := NewQueue[int]()
	su.Zero(e.Dequeue())
}

func (su *QueueSuite) Test_Count_Good() {
	q := su.mockQueue()
	su.Equal(len(q.data), q.Len())

	e := NewQueue[int]()
	su.Equal(0, e.Len())
}

func (su *QueueSuite) Test_Peek_Good() {
	q := su.mockQueue()
	su.Equal(-1, q.Peek())
	su.Equal(-1, q.data[0])

	e := NewQueue[int]()
	su.Zero(e.Peek())
}

func (su *QueueSuite) Test_ToSlice_Good() {
	q := su.mockQueue()
	sli := q.ToSlice()
	su.Require().Equal(len(q.data), len(sli))
	for i := range sli {
		su.Equal(q.data[i], sli[i])
	}

	sli[0] = 123
	su.Equal(-1, q.data[0])
	su.Equal(123, sli[0])
}

func (su *QueueSuite) Test_Clear_Good() {
	q := su.mockQueue()
	q.Clear()
	for i := range q.data {
		su.Zero(q.data[i])
	}
}

func (su *QueueSuite) Test_Shrink_Good() {
	q := su.mockQueue()
	q.Shrink(100)
	su.Equal(10, len(q.data))
	q.Shrink(10)
	su.Equal(10, len(q.data))
	q.Shrink(5)
	su.Equal(5, len(q.data))
	q.Shrink(0)
	su.Equal(0, len(q.data))
}
