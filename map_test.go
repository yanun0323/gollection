package gollection

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMap(t *testing.T) {
	suite.Run(t, new(MapSuite))
}

type MapSuite struct {
	suite.Suite
}

func (su *MapSuite) Test_SyncMap() {
	m := sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)
	v, loaded := m.LoadAndDelete(1)
	fmt.Println("LoadAndDelete", v, loaded)

	v, loaded = m.LoadOrStore(4, 0)
	fmt.Println("LoadOrStore", v, loaded)
	v, loaded = m.LoadOrStore(4, 1)
	fmt.Println("LoadOrStore", v, loaded)
	v, loaded = m.LoadOrStore(4, 2)
	fmt.Println("LoadOrStore", v, loaded)

	v, loaded = m.Swap(6, 0)
	fmt.Println("Swap", v, loaded)

	v, loaded = m.Swap(6, 6)
	fmt.Println("Swap", v, loaded)
}
