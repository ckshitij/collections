package set

type Set[T comparable] struct {
	table map[T]struct{}
}

// NewSet creates and returns a new Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		table: make(map[T]struct{}),
	}
}

// Add inserts an element into the Set.
func (s *Set[T]) Add(value T) {
	s.table[value] = struct{}{}
}

// Remove deletes an element from the Set.
func (s *Set[T]) Remove(value T) {
	delete(s.table, value)
}

// Contains checks if an element exists in the Set.
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.table[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.table)
}
