package heap

// Heap represents a generic heap structure.
// T is the type of the elements stored in the heap, and the compare function defines the heap property.
type Heap[T any] struct {
	table   []T           // The underlying table that holds heap elements.
	compare Comparable[T] // A comparison function to define heap order.
}

// Comparable defines a type for the comparison function used to maintain the heap property.
// The function returns true if the first argument should come before the second.
type Comparable[T any] func(a T, b T) bool

// NewHeap creates a new Heap with the provided comparison function and optional initial elements.
// It constructs the heap by applying the comparison function.
// Elements can be passed variadically to be added initially to the heap.
func NewHeap[T any](compFunc Comparable[T], elements ...T) Heap[T] {
	table := []T{}
	table = append(table, elements...) // Initialize the table with the provided elements.
	hp := Heap[T]{
		table:   table,
		compare: compFunc,
	}
	hp.buildHeap() // Build the heap to satisfy the heap property.
	return hp
}

// Push adds a new element to the heap and restores the heap property by sifting the element up.
func (hp *Heap[T]) Push(data T) {
	hp.table = append(hp.table, data) // Append the new element to the end of the heap.
	hp.siftUp(hp.Size() - 1)          // Restore heap property by sifting up.
}

// left returns the index of the left child of the element at the given index.
func (_ *Heap[T]) left(index int) int {
	return (2 * index) + 1
}

// right returns the index of the right child of the element at the given index.
func (_ *Heap[T]) right(index int) int {
	return (2 * index) + 2
}

// parent returns the index of the parent of the element at the given index.
func (_ *Heap[T]) parent(index int) int {
	return (index - 1) / 2
}

// Pop removes and returns the root element (the smallest/largest element depending on the heap type).
// The heap property is restored after removing the root.
func (hp *Heap[T]) Pop() T {
	var val T
	if hp.Size() <= 0 {
		return val // Return zero value if heap is empty.
	}
	val = hp.table[0]                 // Get the root element.
	hp.swap(0, hp.Size()-1)           // Swap root with the last element.
	hp.table = hp.table[:hp.Size()-1] // Remove the last element.
	hp.heapify(0)                     // Restore heap property by heapifying from the root.
	return val
}

// isValidInd checks if the provided index is within the valid range of the heap.
func (hp *Heap[T]) isValidInd(ind int) bool {
	return ind >= 0 && ind < hp.Size()
}

// swap exchanges the elements at the two provided indices in the heap.
func (hp *Heap[T]) swap(a, b int) {
	hp.table[a], hp.table[b] = hp.table[b], hp.table[a]
}

// heapify restores the heap property from the element at the given index downwards.
// It compares the element with its children and swaps it with the appropriate child if necessary.
func (hp *Heap[T]) heapify(rootInd int) {
	left := hp.left(rootInd)
	right := hp.right(rootInd)
	cmpInd := rootInd

	// Compare the left child.
	if hp.isValidInd(left) && hp.compare(hp.table[left], hp.table[cmpInd]) {
		cmpInd = left
	}

	// Compare the right child.
	if hp.isValidInd(right) && hp.compare(hp.table[right], hp.table[cmpInd]) {
		cmpInd = right
	}

	// If root is not in the correct place, swap and heapify further down.
	if cmpInd != rootInd {
		hp.swap(rootInd, cmpInd)
		hp.heapify(cmpInd)
	}
}

// buildHeap converts the current list of elements into a valid heap by calling heapify from the middle of the heap to the root.
func (hp *Heap[T]) buildHeap() {
	totalElement := hp.Size()
	for i := totalElement / 2; i >= 0; i-- {
		hp.heapify(i)
	}
}

// siftUp restores the heap property by moving the element at the given index upwards
// until it is in the correct position relative to its parent.
func (hp *Heap[T]) siftUp(ind int) {
	parent := hp.parent(ind)
	if ind > 0 && hp.compare(hp.table[ind], hp.table[parent]) {
		hp.swap(ind, parent)
		hp.siftUp(parent)
	}
}

// Size returns the current number of elements in the heap.
func (hp *Heap[T]) Size() int {
	return len(hp.table)
}
