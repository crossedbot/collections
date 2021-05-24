package stack

import ()

type Stack interface {
	// Len returns the length of the stack.
	Len() int

	// Pop returns an item on top of the stack.
	Pop() interface{}

	// Push puts the item on top of the stack.
	Push(value interface{})
}

type stack struct {
	data []interface{}
}

func New() Stack {
	return &stack{}
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}
	item := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return item
}

func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
}
