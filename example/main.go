package main

import (
	"github.com/charmbracelet/log"
	"github.com/prog-lang/pure/machine"
	. "github.com/prog-lang/pure/opcode"
	"github.com/prog-lang/pure/std"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

const (
	start = 0
	zero  = 45
	life  = 55
)

var (
	data = []uint8("https://google.com")
	code = []uint8{
		// $start: () -> ()
		PUSH_FN, uint8(std.Print), 0, 0, 0, // PUSH print
		PUSH_CMD, zero, 0, 0, 0, // PUSH $zero
		PUSH_CMD, life, 0, 0, 0, // PUSH $life
		PUSH_BOOL, 1, 0, 0, 0, // PUSH true
		BRANCH, 0, 0, 0, 0, // IF true -> $life
		CALL, 0, 0, 0, 0, // CALL $life
		FEED, 1, 0, 0, 0, // FEED 2
		CALL, 0, 0, 0, 0, // CALL print
		RETURN, 0, 0, 0, 0,

		// $zero: () -> i32
		PUSH_I32, 0, 0, 0, 0, // PUSH 0
		RETURN, 0, 0, 0, 0,

		// $life: () -> i32
		PUSH_FN, uint8(std.Add_I32), 0, 0, 0, // PUSH add[i32]
		PUSH_I32, 2, 0, 0, 0, // PUSH 2
		PUSH_I32, 40, 0, 0, 0, // PUSH 40
		FEED, 2, 0, 0, 0, // FEED 2
		CALL, 0, 0, 0, 0, // CALL add[i32]
		RETURN, 0, 0, 0, 0,
	}
)

func main() {
	machine.MakeCmd(
		machine.NewROM(data),
		machine.NewROM(code),
		start,
	).Call()
}
