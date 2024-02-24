package machine

// Magic constants.
const (
	SizeU8  = 1
	SizeI32 = 4

	SizeOpcode      = SizeU8
	SizeOperand     = SizeI32
	SizeInstruction = SizeOpcode + SizeOperand
)
