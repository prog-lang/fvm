package opcode

const (
	NOP uint32 = iota // DO NOTHING

	/* Stack manipulation */

	PUSH_UNIT // Push Unit onto the stack
	PUSH_BOOL // Push Bool onto the stack
	PUSH_U8   // Push U8 onto the stack
	PUSH_I32  // Push I32 onto the stack
	PUSH_FN   // Push Fn onto the stack (std)
	PUSH_PROC // Push Proc onto the stack
	PUSH_ARG  // Push Proc argument (by its index) onto the stack
	DROP      // Drop top value off of the stack

	/* Program flow */

	FEED   // Feed N top values into the function beneath
	EXEC   // Execute the Cmd (= PUSH_UNIT; FEED 1;)
	RETURN // Return from the Proc

	Count
)
