package machine

// Byte sizes for different things.
const (
	SizeU8  = 1
	SizeI32 = 4 * SizeU8
	SizeU32 = 4 * SizeU8
	SizeI64 = 8 * SizeU8
	SizeU64 = 8 * SizeU8

	SizeOpcode      = SizeU64
	SizeOperand     = SizeU64
	SizeInstruction = SizeOpcode + SizeOperand
)
