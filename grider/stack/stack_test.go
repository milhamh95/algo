package stack

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestStack_Size(t *testing.T) {
	is := is.New(t)
	curStack := NewStack()
	curStack.items = []int{1}

	res := curStack.size()
	is.Equal(1, res)
}

func TestStack_Search(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{1}

		res := curStack.search(1)
		is.Equal(true, res)
	})

	t.Run("not found", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{1}

		res := curStack.search(2)
		is.Equal(false, res)
	})
}

func TestStack_Push(t *testing.T) {
	is := is.New(t)
	curStack := NewStack()
	curStack.items = []int{}

	curStack.push(1)
	is.Equal(1, curStack.size())
}

func TestStack_Pop(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{1}

		item, err := curStack.pop()
		is.Equal(1, item)
		is.NoErr(err)
		is.Equal(0, curStack.size())
	})

	t.Run("pop empty stack", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{}

		item, err := curStack.pop()
		is.Equal(0, item)
		is.Equal(errors.New("there is no item in stack"), err)
		is.Equal(0, curStack.size())
	})
}

func TestStack_Peek(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{1}

		item, err := curStack.peek()
		is.Equal(1, item)
		is.NoErr(err)
		is.Equal(1, curStack.size())
	})

	t.Run("peek empty stack", func(t *testing.T) {
		curStack := NewStack()
		curStack.items = []int{}

		item, err := curStack.peek()
		is.Equal(0, item)
		is.Equal(errors.New("there is no item in stack"), err)
		is.Equal(0, curStack.size())
	})
}
