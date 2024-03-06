package machine

import (
	"bytes"
	"testing"

	. "github.com/prog-lang/pure/opcode"
	"github.com/prog-lang/pure/std"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	var data []uint8
	dataLength := U64AsU8x8(uint64(len(data)))
	dataSection := append(dataLength, data...)

	first := addr(8)
	codeSection := []uint8{
		// main :: Str
		uint8(PUSH_FN), 0, 0, 0, uint8(std.Print), 0, 0, 0, // print
		uint8(PUSH_FN), 0, 0, 0, uint8(std.Show_I32), 0, 0, 0, // print show[i32]
		uint8(PUSH_PROC), 0, 0, 0, first, 0, 0, 0, // print show[i32] first
		uint8(PUSH_I32), 0, 0, 0, 42, 0, 0, 0, // print show[i32] first 42
		uint8(FEED), 0, 0, 0, 1, 0, 0, 0, // print show[i32] 42
		uint8(FEED), 0, 0, 0, 1, 0, 0, 0, // print "42"
		uint8(FEED), 0, 0, 0, 1, 0, 0, 0, // Cmd Str
		uint8(RETURN), 0, 0, 0, 0, 0, 0, 0,
		// first :: I32 -> I32
		uint8(NOP), 0, 0, 0, 1, 0, 0, 0,
		uint8(PUSH_ARG), 0, 0, 0, 0, 0, 0, 0,
		uint8(RETURN), 0, 0, 0, 0, 0, 0, 0,
	}

	r := bytes.NewReader(append(dataSection, codeSection...))
	proc, err := SourceFromReader(r).Main()
	assert.NoError(t, err)
	putStr, ok := proc.Exec().(Cmd)
	assert.True(t, ok)
	assert.Equal(t, "42", putStr())
}

func addr(index uint8) uint8 {
	return index * 8
}
