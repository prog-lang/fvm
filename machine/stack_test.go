package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int](0, 1).Push(2, 3, 4, 5)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, s.values)
	assert.Equal(t, 5, s.Pop())
	assert.Equal(t, []int{0, 1, 2, 3, 4}, s.values)
	assert.Equal(t, []int{2, 3, 4}, s.Take(3))
	assert.Equal(t, []int{0, 1}, s.values)
	assert.Equal(t, 1, s.Peek())
	assert.Equal(t, []int{0, 1}, s.values)
	assert.Equal(t, []int{0, 1}, s.Glance(2))
	assert.Equal(t, []int{0, 1}, s.values)
	s.Drop(2)
	assert.True(t, s.Empty())
}
