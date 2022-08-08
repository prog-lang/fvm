package main

import (
	. "machine"
)

// Program labels are ROM addresses of specific commands that we can jump to
// using the JUMP and CALL instructions.
const (
	start int32 = 4
	magic int32 = 16
)

func main() {
	var data []byte

	program := []int32{
		CALL, start, // @setup
		EXIT, 0,
		PUSH, 2, // @start
		PUSH, 1,
		SUB, 0,
		COND, 0,
		BR, magic,
		DONE, 0,
		PUSH, 42, // @magic
		DONE, 0,
		EXIT, 0, // catch all exit
	}

	New(data, program).Run()
}
