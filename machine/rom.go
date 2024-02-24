package machine

type ROM struct {
	bytes []uint8
}

func NewROM(bytes []uint8) *ROM {
	return &ROM{
		bytes: bytes,
	}
}

func (rom *ROM) ReadAt(addr, length int32) []uint8 {
	return rom.bytes[addr : addr+length]
}

func (rom *ROM) Fetch(addr int32) Do {
	opcode := rom.u32(addr)
	operand := rom.u8x4(addr + SizeOpcode)
	return is[opcode](operand)
}

func (rom *ROM) u32(addr int32) uint32 {
	return U8x4AsU32(rom.u8x4(addr))
}

func (rom *ROM) u8x4(addr int32) []uint8 {
	return rom.ReadAt(addr, 4)
}
