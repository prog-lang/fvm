package machine

import "fmt"

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		values: make([]T, 0, 100),
	}
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (value T) {
	value = s.Peek()
	s.Drop()
	return
}

func (s *Stack[T]) Peek() T {
	return s.values[s.Top()]
}

func (s *Stack[T]) Drop() {
	s.values = s.values[:s.Top()]
}

func (s *Stack[T]) Top() int {
	return len(s.values) - 1
}

func (s *Stack[T]) String() string {
	return fmt.Sprint("$", s.values)
}
