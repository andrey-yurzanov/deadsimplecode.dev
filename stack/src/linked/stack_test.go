package stack_test

import (
	"testing"

	stack "deadsimplecode.dev/stack/linked"
)

func TestSuccess(t *testing.T) {
	stack := &stack.Stack[int]{}
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	for i := 99; i > 0; i-- {
		if i != stack.Pop() {
			t.Error("Value not equals")
		}
	}
}

func TestFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic is expected")
		}
	}()

	stack := &stack.Stack[int]{}
	stack.Pop()
}
