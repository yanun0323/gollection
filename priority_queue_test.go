package gollection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPriorityQueue(t *testing.T) {
	suite.Run(t, new(PriorityQueueSuite))
}

type PriorityQueueSuite struct {
	suite.Suite
}

func (su *PriorityQueueSuite) SetupSuite() {
}

func (su *PriorityQueueSuite) SetupTest() {

}

func (su *PriorityQueueSuite) Test() {
	pq := NewPriorityQueue[int](func(i1, i2 int) bool {
		return i1 > i2
	})

	pq.Enqueue(7, 3, 8, 2, 6, 1, 4)
	su.Equal(7, pq.Len())

	su.Equal(8, pq.Dequeue())
	su.Equal(6, pq.Len())

	su.IsDecreasing(pq.ToSlice(), pq.ToSlice())

	su.Equal(7, pq.Peek())
	su.Equal(6, pq.Len())
}
