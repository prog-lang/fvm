package machine

const (
	/* SETTINGS */

	defaultDataStackCapacity = 1 << 10
)

const (
	/* MAGIC CONSTANTS */

	SizeU8  = 1
	SizeI32 = 4

	SizeOpcode      = SizeU8
	SizeOperand     = SizeI32
	SizeInstruction = SizeOpcode + SizeOperand
)
