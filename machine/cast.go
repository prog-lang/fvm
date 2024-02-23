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

func I32AsBool(i int32) bool {
	return i != 0
}

func U8AsBool(b uint8) bool {
	return b != 0
}
