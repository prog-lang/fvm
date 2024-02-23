package machine

import (
	"machine/std"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFn(t *testing.T) {
	add0 := MakeFn(stdlib[std.Add_I32])
	add1 := add0.Feed([]Object{1})
	add2 := add1.Feed([]Object{2})
	assert.Len(t, add0.args, 0)
	assert.Len(t, add1.(Fn).args, 1)
	assert.Len(t, add2.(Fn).args, 2)
}
