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
	mockQueue func() queue
}

func (su *QueueSuite) SetupTest() {
	su.mockQueue = func() queue {
		return queue{
			data: []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
	}
}

func (su *QueueSuite) Test_NewQueue_Good() {
	q := NewQueue()
	su.Equal(0, len(q.data))
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
	su.Equal(0, q.Dequeue())
	su.Equal(1, q.data[0])
	su.Equal(9, len(q.data))

	e := queue{}
	su.Nil(e.Dequeue())
}

func (su *QueueSuite) Test_Count_Good() {
	q := su.mockQueue()
	su.Equal(len(q.data), q.Count())

	e := queue{}
	su.Equal(0, e.Count())
}

func (su *QueueSuite) Test_Peek_Good() {
	q := su.mockQueue()
	su.Equal(0, q.Peek())
	su.Equal(0, q.data[0])

	e := queue{}
	su.Nil(e.Peek())
}

func (su *QueueSuite) Test_ToSlice_Good() {
	q := su.mockQueue()
	sli := q.ToSlice()
	su.Require().Equal(len(q.data), len(sli))
	for i := range sli {
		su.Equal(q.data[i], sli[i])
	}

	sli[0] = 123
	su.Equal(0, q.data[0])
	su.Equal(123, sli[0])
}
