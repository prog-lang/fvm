package opcode

const (
	NOP byte = iota // DO NOTHING

	/* Stack manipulation */

	BOOL_PUSH // Push bool on the stack
	U8_PUSH   // Push u8 on the stack
	I32_PUSH  // Push i32 on the stack
	DROP      // Drop top value on the stack

	/* Program flow */

	JUMP // JUMP IP to the specified instruction address in ROM
	CALL // CALL pushes return address onto CS and then jumps
	BR   // BR will perform a CALL if top DS value is true
	DONE // DONE returns from the routine

	Count
)
