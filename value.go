package gollection

type Value[T any] interface {
	CompareAndSwap(old T, new T) (swapped bool)
	Load() (val T)
	TryLoad() (val T, loaded bool)
	Store(val T)
	Swap(new T) (old T)
}
