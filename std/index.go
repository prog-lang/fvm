package std

const (
	ID uint32 = iota // id :: Va. a -> a
	If               // if :: Va. Bool -> a -> a -> a

	/* I32 operations */

	Add_I32  // add :: I32 -> I32 -> I32
	Sub_I32  // sub :: I32 -> I32 -> I32
	Mul_I32  // mul :: I32 -> I32 -> I32
	Div_I32  // div :: I32 -> I32 -> I32
	Show_I32 // show :: I32 -> Str

	/* Cmd operations */

	Cmd_Cmd // cmd :: Va. a -> Cmd a
	// Async_Cmd // async :: Va. Cmd a -> Cmd Unit
	Map_Cmd   // map :: Va, b. (a -> b) -> Cmd a -> Cmd b
	Swap_Cmd  // swap :: Va, b. a -> Cmd b -> Cmd a
	Then_Cmd  // then :: Va, b. Cmd a -> Cmd b -> Cmd b
	Chain_Cmd // chain :: Va, b. Cmd a -> (a -> Cmd b) -> Cmd b

	/* I/O operations */

	Print_Str // print :: Str -> Cmd Str

	Count
)
