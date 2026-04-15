package stack

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	v := s.data[len(s.data)-1]
	var zero T
	s.data[len(s.data)-1] = zero
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
