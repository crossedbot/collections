package randomstack

import (
	"crypto/rand"
	"math/big"
)

// RandomStack is an interface to a randomly accessed "stack" collection.
type RandomStack interface {
	// Len returns the length of the stack.
	Len() int

	// Pop returns a random item from the stack.
	Pop() interface{}

	// Push puts the item in the stack.
	Push(value interface{})
}

// randomstack represents a randomly accessed stack.
type randomstack struct {
	data []interface{}
}

// New returns a new RandomStack.
func New() RandomStack {
	return &randomstack{}
}

// Len returns the length of the stack.
func (rs *randomstack) Len() int {
	return len(rs.data)
}

// Pop returns a random item from the stack.
func (rs *randomstack) Pop() interface{} {
	if rs.Len() == 0 {
		return nil
	}
	i_, _ := rand.Int(rand.Reader, big.NewInt(int64(rs.Len())))
	i := int(i_.Int64())
	item := rs.data[i]
	rs.remove(i)
	return item
}

// Push puts the item in the stack.
func (rs *randomstack) Push(value interface{}) {
	rs.data = append(rs.data, value)
}

// remove removes the item at index 'i' from the stack.
func (rs *randomstack) remove(i int) {
	if i >= 0 && i < rs.Len() {
		// since this "stack" is random, disregard order for efficiency
		rs.data[i] = rs.data[rs.Len()-1]
		rs.data = rs.data[:rs.Len()-1]
	}
}
