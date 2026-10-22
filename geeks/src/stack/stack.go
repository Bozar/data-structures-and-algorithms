package stack

import ()

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) < 1
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Push(data T) {
	s.elements = append(s.elements, data)
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var null T
		return null, false
	}

	elements := s.elements
	idx := len(elements) - 1
	return elements[idx], true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var null T
		return null, false
	}

	var null T
	idx := len(s.elements) - 1
	data := s.elements[idx]
	s.elements[idx] = null
	s.elements = s.elements[:idx]
	return data, true
}
