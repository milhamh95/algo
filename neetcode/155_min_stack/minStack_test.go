package minstack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	newStack := &MinStack{}

	tests := []struct {
		name string
		want *MinStack
	}{
		{
			name: "init constructor not return nil",
			want: newStack,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)

	stack.Print()
	top := stack.Top()
	fmt.Println("top:", top)

	stack.Pop()
	stack.Print()

	stack.Push(3)
	stack.Push(5)

	min := stack.GetMin()
	fmt.Println("min:", min)
}
