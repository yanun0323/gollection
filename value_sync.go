package gollection

import "sync/atomic"

type syncValue[T any] struct {
	val atomic.Value
}

func NewSyncValue[T any](val ...T) Value[T] {
	v := atomic.Value{}
	if len(val) != 0 {
		v.Store(val[0])
	}

	return &syncValue[T]{
		val: v,
	}
}

// CompareAndSwap executes the compare-and-swap operation for the Value.
func (v *syncValue[T]) CompareAndSwap(old T, new T) (swapped bool) {
	return v.val.CompareAndSwap(old, new)
}

// Load returns the value set by the most recent Store.
// It returns zero value if there has been no call to Store for this Value.
func (v *syncValue[T]) Load() (val T) {
	if vv, ok := v.val.Load().(T); ok {
		return vv
	}

	var zero T
	return zero
}

// Store sets the value of the Value v to val.
func (v *syncValue[T]) Store(val T) {
	v.val.Store(val)
}

// Swap stores new into Value and returns the previous value. It returns zero value if
// the Value is empty.
func (v *syncValue[T]) Swap(new T) (old T) {
	if vv, ok := v.val.Swap(new).(T); ok {
		return vv
	}

	var zero T
	return zero
}
