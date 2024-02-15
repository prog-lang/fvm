package opcode

const (
	/* Stack manipulation */

	PUSH int32 = iota // PUSH int32 on the stack
	DROP              // DROP top value on the stack

	/* RAM manipulation */
	STORE // STORE top stack value at some location in RAM
	LOAD  // LOAD value from RAM onto a stack

	/* Program flow */

	JUMP // JUMP IP to the specified instruction address in ROM
	CALL // CALL pushes return address onto CS and then jumps
	BR   // BR will perform a CALL if top DS value is true
	DONE // DONE jumps back to the calling subroutine
	EXIT // EXIT the program

	Count
)
