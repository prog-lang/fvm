package machine

const (
	/* SETTINGS */

	defaultRAMCapacity       = 1 << 20
	defaultDataStackCapacity = 1 << 10
	defaultCallStackCapacity = 1 << 8
)

const (
	/* MAGIC CONSTANTS */

	U8Size  = 1
	I32Size = 4

	OpcodeSize      = U8Size
	OperandSize     = I32Size
	InstructionSize = OpcodeSize + OperandSize
)
