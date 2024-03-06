package machine

import (
	"testing"

	"github.com/prog-lang/pure/std"
	"github.com/stretchr/testify/assert"
)

func TestFn(t *testing.T) {
	add0 := stdlib[std.Add_I32]
	add1 := add0.Apply(int32(1))
	result := add1.(Fn).Apply(int32(2))
	assert.Len(t, add0.args, 0)
	assert.Len(t, add1.(Fn).args, 1)
	assert.Equal(t, int32(3), result.(int32))
}
