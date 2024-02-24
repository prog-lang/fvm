package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestU8x4AsI32(t *testing.T) {
	assert.Equal(t, int32(1), U8x4AsI32([]uint8{1, 0, 0, 0}))
}

func TestU8x8AsI64(t *testing.T) {
	assert.Equal(t, uint64(298), U8x8AsU64([]uint8{42, 1, 0, 0, 0, 0, 0, 0}))
}
