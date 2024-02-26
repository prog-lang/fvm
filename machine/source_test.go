package machine

import (
	"bytes"
	"testing"

	"github.com/charmbracelet/log"
	. "github.com/prog-lang/pure/opcode"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var data []uint8
	dataLength := U64AsU8x8(uint64(len(data)))
	dataSection := append(dataLength, data...)

	const first = 32
	codeSection := []uint8{
		uint8(PUSH_CMD), 0, 0, 0, first, 0, 0, 0,
		uint8(PUSH_I32), 0, 0, 0, 42, 0, 0, 0,
		uint8(FEED), 0, 0, 0, 1, 0, 0, 0,
		uint8(RETURN), 0, 0, 0, 0, 0, 0, 0,
		uint8(NOP), 0, 0, 0, 1, 0, 0, 0, // first : i32 -> i32
		uint8(PUSH_ARG), 0, 0, 0, 0, 0, 0, 0,
		uint8(RETURN), 0, 0, 0, 0, 0, 0, 0,
	}

	r := bytes.NewReader(append(dataSection, codeSection...))
	cmd, err := SourceFromReader(r).Main()
	assert.NoError(t, err)
	assert.Equal(t, int32(42), cmd.Feed(Unit{}).(int32))
}
