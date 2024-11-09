package gollection

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSyncValue(t *testing.T) {
	suite.Run(t, new(SyncValueSuite))
}

type SyncValueSuite struct {
	suite.Suite
	ctx context.Context
}

func (su *SyncValueSuite) SetupSuite() {
	su.ctx = context.Background()
}

func (su *SyncValueSuite) Test() {
	val := NewSyncValue[func()]()
	su.Nil(val.Load())
	su.NotPanics(func() {
		val.Store(nil)
	})
	su.Nil(val.Load())
	val.Store(func() {})
	su.NotNil(val.Load())

	su.NotNil(val.Swap(nil))
	su.Nil(val.Swap(nil))
}
