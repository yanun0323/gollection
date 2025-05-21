package gollection

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSyncQueue(t *testing.T) {
	suite.Run(t, new(SyncQueueSuite))
}

type SyncQueueSuite struct {
	suite.Suite
	mockQueue func() syncQueue[int]
}

func (su *SyncQueueSuite) SetupTest() {
	su.mockQueue = func() syncQueue[int] {
		return syncQueue[int]{
			rwLock: &sync.RWMutex{},
			q: queue[int]{
				data: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}
	}
}

func (su *SyncQueueSuite) Test_NewQueue_Good() {
	q := NewQueue[int]()
	su.NotNil(q)
}

func (su *SyncQueueSuite) Test_Enqueue_Good() {
	q := su.mockQueue()
	q.Enqueue(10)
	su.Equal(10, q.q.data[10])
	q.Enqueue(11, 12, 13)
	su.Equal(11, q.q.data[11])
	su.Equal(12, q.q.data[12])
	su.Equal(13, q.q.data[13])
}

func (su *SyncQueueSuite) Test_Dequeue_Good() {
	q := su.mockQueue()
	su.Equal(-1, q.Dequeue())
	su.Equal(1, q.q.data[0])
	su.Equal(9, len(q.q.data))

	e := NewQueue[int]()
	su.Zero(e.Dequeue())
}

func (su *SyncQueueSuite) Test_Count_Good() {
	q := su.mockQueue()
	su.Equal(len(q.q.data), q.Len())

	e := NewQueue[int]()
	su.Equal(0, e.Len())
}

func (su *SyncQueueSuite) Test_Peek_Good() {
	q := su.mockQueue()
	su.Equal(-1, q.Peek())
	su.Equal(-1, q.q.data[0])

	e := NewQueue[int]()
	su.Zero(e.Peek())
}

func (su *SyncQueueSuite) Test_ToSlice_Good() {
	q := su.mockQueue()
	sli := q.ToSlice()
	su.Require().Equal(len(q.q.data), len(sli))
	for i := range sli {
		su.Equal(q.q.data[i], sli[i])
	}

	sli[0] = 123
	su.Equal(-1, q.q.data[0])
	su.Equal(123, sli[0])
}

func (su *SyncQueueSuite) Test_Clear_Good() {
	q := su.mockQueue()
	q.Clear()
	for i := range q.q.data {
		su.Zero(q.q.data[i])
	}
}

func (su *SyncQueueSuite) Test_Shrink_Good() {
	q := su.mockQueue()
	q.Shrink(100)
	su.Equal(10, len(q.q.data))
	q.Shrink(10)
	su.Equal(10, len(q.q.data))
	q.Shrink(5)
	su.Equal(5, len(q.q.data))
	q.Shrink(0)
	su.Equal(0, len(q.q.data))
}
