package stack

import "errors"

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) size() int {
	return len(s.items)
}

func (s *Stack) search(item int) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}

	return false
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("there is no item in stack")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("there is no item in stack")
	}

	lenItem := len(s.items)
	item := s.items[lenItem-1]

	s.items = s.items[:lenItem-1]
	return item, nil
}
