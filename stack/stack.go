package stack

import (
	"errors"

	"github.com/ckshitij/collections/list"
)

// Stack represents a generic stack data structure that holds elements of any type.
type Stack[T any] struct {
	head *list.List[T]
}

// NewStack creates and returns a new instance of Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: list.NewList[T](),
	}
}

// Push adds a new element to the top of the stack.
func (st *Stack[T]) Push(value T) {
	st.head.PushFront(value)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (st *Stack[T]) IsEmpty() bool {
	return st.head.Len() < 1
}

// Pop removes the top element from the stack.
// Returns an error if the stack is empty.
func (st *Stack[T]) Pop() error {
	if st.IsEmpty() {
		return errors.New("invalid operation: empty stack")
	}
	st.head.PopFront()
	return nil
}

// Top returns the top element of the stack without removing it.
// Returns the zero value of the type if the stack is empty.
func (st *Stack[T]) Top() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Front().Element()
	}
	return empty
}
