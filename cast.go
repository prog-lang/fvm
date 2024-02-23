package machine

func I32AsU8x4(i int32) (b []byte) {
	b = make([]byte, 4)
	for x := 0; x < 4; x++ {
		b[x] = byte(i >> (8 * x))
	}
	return
}

func U8x4AsI32(b []byte) (i int32) {
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

func BoolAsU8(t bool) byte {
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

func U8AsBool(b byte) bool {
	return b != 0
}
