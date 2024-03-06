package std

const (
	ID uint32 = iota // id :: Va. a -> a

	Add_I32  // add[i32] :: I32 -> I32 -> I32
	Sub_I32  // sub[i32] :: I32 -> I32 -> I32
	Mul_I32  // mul[i32] :: I32 -> I32 -> I32
	Div_I32  // div[i32] :: I32 -> I32 -> I32
	Show_I32 // show[i32] :: I32 -> Str

	Print // print :: Str -> Cmd Str

	Count
)
