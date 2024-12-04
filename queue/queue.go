package queue

import (
	"errors"

	"github.com/ckshitij/collections/list"
)

// Queue represents a generic queue data structure that holds elements of any type.
type Queue[T any] struct {
	head *list.List[T]
}

// NewQueue creates and returns a new instance of Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: list.NewList[T](),
	}
}

// Enqueue adds a new element to the back of the queue.
func (st *Queue[T]) Enqueue(value T) {
	st.head.PushBack(value)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (st *Queue[T]) IsEmpty() bool {
	return st.head.Len() < 1
}

// Dequeue removes the front element from the queue.
// Returns an error if the queue is empty.
func (st *Queue[T]) Dequeue() error {
	if st.IsEmpty() {
		return errors.New("invalid operation: empty queue")
	}
	st.head.PopFront()
	return nil
}

// Front returns the front element of the queue without removing it.
// Returns the zero value of the type if the queue is empty.
func (st *Queue[T]) Front() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Front().Element()
	}
	return empty
}

// Back returns the back element of the queue without removing it.
// Returns the zero value of the type if the queue is empty.
func (st *Queue[T]) Back() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Back().Element()
	}
	return empty
}
