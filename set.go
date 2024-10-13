package gollection

// Set is an interface for a set data structure.
type Set[T comparable] interface {
	// Contain returns true if the set contains the element.
	Contain(T) bool

	// Len returns the number of elements in the set.
	Len() int

	// Insert adds elements to the set.
	Insert(...T)

	// Remove removes all elements in a from the set.
	Remove(...T)

	// ToSlice returns a copy of the set as a slice.
	ToSlice() []T
}

type set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a new set.
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
