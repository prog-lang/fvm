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

func (rom *ROM) ReadAt(addr, length uint64) []uint8 {
	return rom.bytes[addr : addr+length]
}

func (rom *ROM) Fetch(addr uint64) Do {
	opcode, operand := rom.FetchInstruction(addr)
	return instructions[opcode](operand)
}

func (rom *ROM) FetchInstruction(addr uint64) (opcode uint64, operand []uint8) {
	return rom.u32(addr), rom.u8x8(addr + SizeOpcode)
}

func (rom *ROM) u32(addr uint64) uint64 {
	return U8x8AsU64(rom.u8x8(addr))
}

func (rom *ROM) u8x8(addr uint64) []uint8 {
	return rom.ReadAt(addr, 8)
}
