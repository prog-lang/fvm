package machine

// Magic constants.
const (
	SizeU8  = 1
	SizeI32 = 4 * SizeU8
	SizeU32 = SizeI32
	SizeI64 = 8 * SizeU8
	SizeU64 = SizeI64

	SizeOpcode      = SizeU32
	SizeOperand     = 4 * SizeU8
	SizeInstruction = SizeOpcode + SizeOperand
)
