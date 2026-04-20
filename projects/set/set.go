package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		result.Add(k)
	}
	for k := range other.m {
		result.Add(k)
	}
	return result
}

func (s *Set[T]) Intersect(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
