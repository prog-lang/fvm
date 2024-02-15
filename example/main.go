package main

import (
	"machine"
	. "machine/opcode"
	"machine/std"

	"github.com/charmbracelet/log"
)

// Program labels are ROM addresses of specific commands that we can jump to
// using the JUMP, BR, and CALL instructions.
const (
	start int32 = 4
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func main() {
	message := "Hello, world!\n"
	length := int32(len(message))
	data := []byte(message)

	rom := []int32{
		// 1. @start is the entrypoint
		// 2. Once @start is DONE => EXIT
		CALL, start, // @setup
		EXIT, 0,

		PUSH, 0, // @start
		PUSH, length,
		CALL, std.Print_data,
		DONE, 0,
	}

	machine.New(data, rom).Run()
}
