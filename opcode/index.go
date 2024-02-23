package opcode

const (
	NOP byte = iota // DO NOTHING

	/* Stack manipulation */

	PUSH_UNIT // Push unit onto the stack
	PUSH_BOOL // Push bool onto the stack
	PUSH_U8   // Push u8 onto the stack
	PUSH_I32  // Push i32 onto the stack
	PUSH_CMD  // Push cmd onto the stack
	DROP      // Drop top value off of the stack

	/* Program flow */

	CALL   // CALL top function off of the stack
	BRANCH // BRANCH left or right based on a condition
	RETURN // Return from the routine

	Count
)
