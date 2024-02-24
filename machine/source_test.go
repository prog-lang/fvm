package machine

import (
	"bytes"
	"testing"

	. "github.com/prog-lang/pure/opcode"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	var data []uint8
	dataLength := U64AsU8x8(uint64(len(data)))
	dataSection := append(dataLength, data...)

	codeSection := []uint8{
		uint8(NOP), 0, 0, 0, 0, 0, 0, 0,
		uint8(PUSH_I32), 0, 0, 0, 42, 0, 0, 0,
		uint8(RETURN), 0, 0, 0, 0, 0, 0, 0,
	}

	r := bytes.NewReader(append(dataSection, codeSection...))
	cmd, err := SourceFromReader(r).MakeCmd()
	assert.NoError(t, err)
	assert.Equal(t, int32(42), cmd.Call().(int32))
}
