package gollection

type Set[T comparable] interface {
	Contain(T) bool
	Len() int
	Insert(...T)
	Remove(...T)
}

type set[T comparable] struct {
	hash map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		hash: make(map[T]bool, 0),
	}
}

func (s *set[T]) Insert(a ...T) {
	for i := range a {
		s.hash[a[i]] = true
	}
}

func (s *set[T]) Remove(a ...T) {
	for i := range a {
		delete(s.hash, a[i])
	}
}

func (s *set[T]) Contain(a T) bool {
	return s.hash[a]
}

func (s *set[T]) Len() int {
	return len(s.hash)
}
