package machine

type ROM struct {
	bytes []byte
}

func NewROM(bytes []byte) *ROM {
	return &ROM{
		bytes: bytes,
	}
}

func (rom *ROM) ReadAt(addr, length int32) []byte {
	return rom.bytes[addr : addr+length]
}

func (rom *ROM) Fetch(addr int32) Do {
	opcode := rom.u8(addr)
	operand := rom.u8x4(addr + OpcodeSize)
	return IS[opcode](operand)
}

func (rom *ROM) u8(addr int32) byte {
	return rom.bytes[addr]
}

func (rom *ROM) u8x4(addr int32) []byte {
	return rom.ReadAt(addr, 4)
}
