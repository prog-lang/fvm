package std

const (
	_ int32 = -iota

	I32_add // func Add(int32, int32) int32
	I32_sub // func Subtract(int32, int32) int32
	I32_mul // func Multiply(int32, int32) int32
	I32_div // func Divide(int32, int32) int32

	I32_byte // func Convert(int32) byte
	Byte_i32 // func Convert(byte) int32

	I32_lt // func LessThan(int32, int32) byte
	I32_eq // func Equal(int32, int32) byte
	I32_gt // func GreaterThan(int32, int32) byte

	Print_data // func Print(addr int32, len int32)

	count
)

const Count = -count
