package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	// Test creating a new set
	s := NewSet[int]()
	assert.NotNil(t, s)
	assert.Equal(t, 0, s.Size())
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()

	// Test adding elements
	s.Add(1)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains(1))

	s.Add(2)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(2))

	// Test adding duplicate element
	s.Add(1)
	assert.Equal(t, 2, s.Size()) // Size should not increase
	assert.True(t, s.Contains(1))
}

func TestRemove(t *testing.T) {
	s := NewSet[int]()

	// Test removing from an empty set
	s.Remove(1) // Should not panic
	assert.Equal(t, 0, s.Size())

	// Test removing existing element
	s.Add(1)
	s.Remove(1)
	assert.Equal(t, 0, s.Size())
	assert.False(t, s.Contains(1))

	// Test removing non-existent element
	s.Add(2)
	s.Remove(3)                  // Non-existent element
	assert.Equal(t, 1, s.Size()) // Size should not change
	assert.True(t, s.Contains(2))
}

func TestContains(t *testing.T) {
	s := NewSet[string]()

	// Test contains on an empty set
	assert.False(t, s.Contains("hello"))

	// Test contains on a set with elements
	s.Add("hello")
	assert.True(t, s.Contains("hello"))
	assert.False(t, s.Contains("world"))

	// Test contains after removal
	s.Remove("hello")
	assert.False(t, s.Contains("hello"))
}

func TestSize(t *testing.T) {
	s := NewSet[float64]()

	// Test size of an empty set
	assert.Equal(t, 0, s.Size())

	// Test size after adding elements
	s.Add(1.1)
	s.Add(2.2)
	assert.Equal(t, 2, s.Size())

	// Test size after adding duplicates
	s.Add(1.1)
	assert.Equal(t, 2, s.Size()) // Size should not increase

	// Test size after removal
	s.Remove(1.1)
	assert.Equal(t, 1, s.Size())
}
