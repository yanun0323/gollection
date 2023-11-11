package gollection

type Set[T comparable] interface {
	Contain(T) bool
	Len() int
	Insert(...T)
	Remove(...T)
	ToSlice() []T
}

type set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		m: make(map[T]struct{}, 0),
	}
}

func (s *set[T]) Insert(a ...T) {
	for i := range a {
		s.m[a[i]] = struct{}{}
	}
}

func (s *set[T]) Remove(a ...T) {
	for i := range a {
		delete(s.m, a[i])
	}
}

func (s *set[T]) Contain(a T) bool {
	_, ok := s.m[a]
	return ok
}

func (s *set[T]) Len() int {
	return len(s.m)
}

func (s *set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.m))
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
