package machine

import "fmt"

type Stack[T any] struct {
	values []T
}

func NewStack[T any](values ...T) *Stack[T] {
	return &Stack[T]{
		values: values,
	}
}

func (s *Stack[T]) Push(values ...T) *Stack[T] {
	s.values = append(s.values, values...)
	return s
}

func (s *Stack[T]) Pop() (creme T) {
	return s.Take(1)[0]
}

func (s *Stack[T]) Take(offset int) (creme []T) {
	creme = s.Glance(offset)
	s.Drop(offset)
	return
}

func (s *Stack[T]) Peek() T {
	return s.Glance(1)[0]
}

func (s *Stack[T]) Glance(offset int) []T {
	return s.values[s.Top(offset):]
}

func (s *Stack[T]) Drop(offset int) *Stack[T] {
	s.values = s.values[:s.Top(offset)]
	return s
}

func (s *Stack[T]) Top(offset int) int {
	return len(s.values) - offset
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func (s *Stack[T]) String() string {
	return fmt.Sprint("$", s.values)
}
