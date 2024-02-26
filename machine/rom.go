package machine

// ROM (a.k.a "Read-Only Memory") implements the Data and Code interfaces. It is
// useful for different operations on bytes.
type ROM struct {
	bytes []uint8
}

func ReadOnly(bytes []uint8) *ROM {
	return &ROM{
		bytes: bytes,
	}
}

func (rom *ROM) ReadAt(addr, length uint32) []uint8 {
	return rom.bytes[addr : addr+length]
}

func (rom *ROM) Fetch(addr uint32) Do {
	opcode, operand := rom.FetchInstruction(addr)
	return instructions[opcode](operand)
}

func (rom *ROM) FetchInstruction(addr uint32) (opcode uint32, operand []uint8) {
	return rom.u32(addr), rom.u8x4(addr + SizeOpcode)
}

func (rom *ROM) u32(addr uint32) uint32 {
	return U8x4AsU32(rom.u8x4(addr))
}

func (rom *ROM) u8x4(addr uint32) []uint8 {
	return rom.ReadAt(addr, 4)
}
