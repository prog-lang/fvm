package main

import (
	. "machine"
)

// Program labels are ROM addresses of specific commands that we can jump to
// using the JUMP, BR, and CALL instructions.
const (
	start int32 = 4
	magic int32 = 16
)

func main() {
	var data []byte

	rom := []int32{
		// 1. @start is the entrypoint
		// 2. Once @start is DONE => EXIT
		CALL, start, // @setup
		EXIT, 0,

		PUSH, 2, // @start
		PUSH, 1,
		SUB, 0,
		I2B, 0,
		BR, magic,
		DONE, 0,

		PUSH, 42, // @magic
		STORE, 0,
		DROP, 4,
		DONE, 0,
	}

	New(data, rom).Run()
}
