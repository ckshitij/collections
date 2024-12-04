package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	// Min-heap comparator
	minHeapComp := func(a, b int) bool {
		return a < b
	}

	// Test creating an empty heap
	h := NewHeap(minHeapComp)
	assert.Equal(t, 0, h.Size(), "Heap size should be zero on creation")

	// Test creating a heap with initial elements
	h = NewHeap(minHeapComp, 5, 3, 8, 1)
	assert.Equal(t, 4, h.Size(), "Heap size should match the number of initial elements")
	assert.Equal(t, 1, h.Pop(), "Heap should pop the smallest element in a min-heap")

	// Test creating a heap with duplicate elements
	h = NewHeap(minHeapComp, 5, 5, 5)
	assert.Equal(t, 3, h.Size(), "Heap size should match the number of initial elements")
	assert.Equal(t, 5, h.Pop(), "Heap should handle duplicate elements correctly")
}

func TestHeapPush(t *testing.T) {
	// Max-heap comparator
	maxHeapComp := func(a, b int) bool {
		return a > b
	}
	h := NewHeap(maxHeapComp)

	// Push elements into the heap
	h.Push(10)
	h.Push(20)
	h.Push(5)

	assert.Equal(t, 3, h.Size(), "Heap size should be 3 after pushing three elements")
	assert.Equal(t, 20, h.Pop(), "Heap should pop the largest element in a max-heap")

	// Push edge case: very large and very small numbers
	h.Push(1000000)
	h.Push(-1000000)
	assert.Equal(t, 1000000, h.Pop(), "Heap should handle very large numbers correctly")
	assert.Equal(t, 10, h.Pop(), "Heap should handle very large numbers correctly")
	assert.Equal(t, 5, h.Pop(), "Heap should handle very large numbers correctly")
	assert.Equal(t, -1000000, h.Pop(), "Heap should handle very small numbers correctly")
}

func TestHeapPop(t *testing.T) {
	// Min-heap comparator
	minHeapComp := func(a, b int) bool {
		return a < b
	}
	h := NewHeap(minHeapComp, 15, 10, 30, 5)

	// Test popping elements from the heap
	assert.Equal(t, 5, h.Pop(), "First popped element should be the smallest")
	assert.Equal(t, 10, h.Pop(), "Second popped element should be the next smallest")
	assert.Equal(t, 15, h.Pop(), "Third popped element should be the next smallest")
	assert.Equal(t, 30, h.Pop(), "Fourth popped element should be the last element")

	// Test popping from an empty heap
	assert.Equal(t, 0, h.Size(), "Heap should be empty after all pops")
	var zeroVal int
	assert.Equal(t, zeroVal, h.Pop(), "Popping from an empty heap should return the zero value of the type")
}

func TestHeapSize(t *testing.T) {
	h := NewHeap(func(a, b int) bool { return a < b })

	assert.Equal(t, 0, h.Size(), "Heap size should initially be zero")

	h.Push(1)
	h.Push(2)
	assert.Equal(t, 2, h.Size(), "Heap size should increase after pushes")

	h.Pop()
	assert.Equal(t, 1, h.Size(), "Heap size should decrease after pops")
}

func TestHeapify(t *testing.T) {
	// Custom comparator for testing
	comp := func(a, b int) bool {
		return a < b
	}

	h := NewHeap(comp, 10, 20, 5)
	assert.Equal(t, 5, h.Pop(), "Heap should be valid after buildHeap")
	assert.Equal(t, 10, h.Pop(), "Heap should still be valid after popping elements")

	// Edge case: Build heap with no elements
	h = NewHeap(comp)
	assert.Equal(t, 0, h.Size(), "Heap should handle empty input during buildHeap")
}

func TestSiftUp(t *testing.T) {
	// Min-heap comparator
	minHeapComp := func(a, b int) bool {
		return a < b
	}
	h := NewHeap(minHeapComp)

	h.Push(20)
	h.Push(10)

	assert.Equal(t, 10, h.Pop(), "Heap property should be restored after siftUp")
	assert.Equal(t, 20, h.Pop(), "Heap property should remain valid after further operations")

	// Edge case: Push into empty heap
	h.Push(50)
	assert.Equal(t, 50, h.Pop(), "Heap should handle siftUp correctly when heap is empty")
}

func TestCustomHeap(t *testing.T) {
	// Custom type
	type customType struct {
		priority int
		name     string
	}

	// Comparator for a max-heap by priority
	comp := func(a, b customType) bool {
		return a.priority > b.priority
	}

	h := NewHeap(comp)
	h.Push(customType{priority: 3, name: "Alice"})
	h.Push(customType{priority: 5, name: "Bob"})
	h.Push(customType{priority: 1, name: "Charlie"})

	top := h.Pop()
	assert.Equal(t, 5, top.priority, "Heap should prioritize elements correctly")
	assert.Equal(t, "Bob", top.name, "Heap should return the element with the correct priority")

	// Edge case: Pop all elements
	h.Pop()
	h.Pop()
	assert.Equal(t, 0, h.Size(), "Heap should be empty after popping all elements")
}

func TestHeapEdgeCases(t *testing.T) {
	// Edge case: Empty heap
	h := NewHeap(func(a, b int) bool { return a < b })
	assert.Equal(t, 0, h.Size(), "Empty heap should have size 0")
	assert.Equal(t, 0, h.Pop(), "Popping from an empty heap should return zero value")

	// Edge case: Single element
	h.Push(42)
	assert.Equal(t, 1, h.Size(), "Heap size should be 1 with a single element")
	assert.Equal(t, 42, h.Pop(), "Popping the single element should return its value")
	assert.Equal(t, 0, h.Size(), "Heap should be empty after popping the single element")

	// Edge case: Duplicate elements
	h.Push(7)
	h.Push(7)
	assert.Equal(t, 2, h.Size(), "Heap should allow duplicate elements")
	assert.Equal(t, 7, h.Pop(), "Heap should handle duplicate elements correctly")
	assert.Equal(t, 7, h.Pop(), "Heap should handle duplicate elements correctly")
	assert.Equal(t, 0, h.Size(), "Heap should be empty after popping duplicates")
}
