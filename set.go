package gollection

type set struct {
	hash map[any]bool
}

func NewSet() set {
	return set{
		hash: make(map[any]bool, 0),
	}
}

func (s *set) Insert(a ...any) {
	for i := range a {
		s.hash[a[i]] = true
	}
}

func (s *set) Remove(a ...any) {
	for i := range a {
		delete(s.hash, a[i])
	}
}

func (s *set) Contain(a any) bool {
	return s.hash[a]
}

func (s *set) Len() int {
	return len(s.hash)
}
