package misc

import "errors"

type StackUtil[T any] struct {
	data []T
}

func (s *StackUtil[T]) Length() int {
	return len(s.data)
}

func (s *StackUtil[T]) Push(item T) {
	if len(s.data) <= 0 {
		s.data = make([]T, 0)
	}

	s.data = append(s.data, item)
}

func (s *StackUtil[T]) Pop() (*T, error) {
	if len(s.data) <= 0 {
		return nil, errors.New("nothing to pop, you doofus")
	}

	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return &ret, nil
}

func (s *StackUtil[T]) Peek() (*T, error) {
	if len(s.data) <= 0 {
		return nil, errors.New("何もない")
	}

	ret := s.data[len(s.data)-1]

	return &ret, nil
}
