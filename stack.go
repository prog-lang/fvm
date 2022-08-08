package machine

import "fmt"

type Stack struct {
	values []int32
}

func NewStack() *Stack {
	return &Stack{
		values: make([]int32, 0, 100),
	}
}

func (s *Stack) Push(value int32) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (value int32) {
	value = s.Peek()
	s.Drop()
	return
}

func (s *Stack) Peek() int32 {
	return s.values[s.Top()]
}

func (s *Stack) Drop() {
	s.values = s.values[:s.Top()]
}

func (s *Stack) Top() int {
	return len(s.values) - 1
}

func (s *Stack) String() string {
	return fmt.Sprint("$", s.values)
}
