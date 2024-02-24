package machine

import (
	"bytes"
	"testing"

	. "github.com/prog-lang/pure/opcode"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	var data []uint8
	dataLength := I32AsU8x4(int32(len(data)))
	dataSection := append(dataLength, data...)

	codeSection := []uint8{
		NOP, 0, 0, 0, 0,
		PUSH_I32, 42, 0, 0, 0,
		RETURN, 0, 0, 0, 0,
	}

	r := bytes.NewReader(append(dataSection, codeSection...))
	cmd, err := SourceFromReader(r).MakeCmd()
	assert.NoError(t, err)
	assert.Equal(t, int32(42), cmd.Call().(int32))
}
