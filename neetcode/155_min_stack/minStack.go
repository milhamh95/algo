package minstack

import (
	"fmt"
	"math"
)

type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *MinStack) Pop() {
	if len(s.data) > 0 {
		s.data[len(s.data)-1] = 0
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}

	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.data) == 0 {
		return 0
	}

	min := math.MaxInt
	for _, v := range s.data {
		if v < min {
			min = v
		}
	}

	return min
}

func (s *MinStack) Print() {
	fmt.Printf("stack: %+v\n", s.data)
	fmt.Printf("len: %v\n", s.data)
	fmt.Printf("cap: %v\n", s.data)
}
