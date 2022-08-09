package machine

func Int32AsBytes(i int32) (b []byte) {
	b = make([]byte, 4)
	for x := 0; x < 4; x++ {
		b[x] = byte(i >> (8 * x))
	}
	return
}

func BytesAsInt32(b []byte) (i int32) {
	for x := 0; x < 4; x++ {
		i |= int32(b[x]) << (8 * x)
	}
	return
}

func BoolAsInt32(t bool) (i int32) {
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

func Int32AsBool(i int32) bool {
	return i != 0
}
