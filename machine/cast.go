package machine

func I32AsU8x4(i int32) (b []uint8) {
	b = make([]uint8, 4)
	for x := 0; x < 4; x++ {
		b[x] = uint8(i >> (8 * x))
	}
	return
}

func U8x4AsI32(b []uint8) (i int32) {
	for x := 0; x < 4; x++ {
		i |= int32(b[x]) << (8 * x)
	}
	return
}

func U8x4AsU32(b []uint8) (i uint32) {
	for x := 0; x < 4; x++ {
		i |= uint32(b[x]) << (8 * x)
	}
	return
}

func U8x8AsI64(b []uint8) (i int64) {
	for x := 0; x < 8; x++ {
		i |= int64(b[x]) << (8 * x)
	}
	return
}

func U8x8AsU64(b []uint8) (i uint64) {
	for x := 0; x < 8; x++ {
		i |= uint64(b[x]) << (8 * x)
	}
	return
}

func U64AsU8x8(u64 uint64) (b []uint8) {
	b = make([]uint8, 8)
	for x := 0; x < 8; x++ {
		b[x] = uint8(u64 >> (8 * x))
	}
	return
}

func BoolAsI32(t bool) int32 {
	if t {
		return 1
	}
	return 0
}

func BoolAsU8(t bool) uint8 {
	if t {
		return 1
	}
	return 0
}

func BoolAsEmoji(t bool) rune {
	if t {
		return '✅'
	}
	return '❌'
}

func I64AsBool(i int64) bool {
	return i != 0
}

func I32AsBool(i int32) bool {
	return i != 0
}

func U8AsBool(b uint8) bool {
	return b != 0
}
