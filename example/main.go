package main

import (
	"fmt"
	"machine"
	. "machine/opcode"

	"github.com/charmbracelet/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

const (
	main_ = 0
	life  = 30
	zero  = 40
)

var (
	data = []byte("https://google.com")
	code = []byte{
		// [$main @0]: () -> i32
		PUSH_CMD, zero, 0, 0, 0, // PUSH $zero
		PUSH_CMD, life, 0, 0, 0, // PUSH $life
		PUSH_BOOL, 1, 0, 0, 0, // PUSH true
		BRANCH, 0, 0, 0, 0, // IF
		CALL, 0, 0, 0, 0, // CALL $life
		RETURN, 0, 0, 0, 0,

		// [$life @30]: () -> i32
		PUSH_I32, 42, 0, 0, 0, // PUSH 42
		RETURN, 0, 0, 0, 0,

		// [$zero @40]: () -> i32
		PUSH_I32, 0, 0, 0, 0, // PUSH 0
		RETURN, 0, 0, 0, 0,
	}
)

func main() {
	result := machine.NewCmd(
		machine.NewROM(data),
		machine.NewROM(code),
		main_,
	).Call()
	fmt.Println(result)
}
