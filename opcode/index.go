package opcode

const (
	NOP uint32 = iota // DO NOTHING

	/* Stack manipulation */

	PUSH_UNIT // Push unit onto the stack
	PUSH_BOOL // Push bool onto the stack
	PUSH_U8   // Push u8 onto the stack
	PUSH_I32  // Push i32 onto the stack
	PUSH_FN   // Push fn onto the stack (std)
	PUSH_CMD  // Push cmd onto the stack
	PUSH_ARG  // Push cmd argument (by its index) onto the stack
	DROP      // Drop top value off of the stack

	/* Program flow */

	FEED   // FEED N top values into the function beneath
	BRANCH // BRANCH left or right based on a condition
	RETURN // Return from the routine

	Count
)
