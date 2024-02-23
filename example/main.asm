start:	NOP				; () -> ()
	PUSH_FN std.Print
	PUSH_CMD @zero
	PUSH_CMD @life
	PUSH_BOOL 1
	BRANCH				; IF true -> $life
	CALL				; CALL $life -> 42
	FEED 1				; std.Print.Feed(42)
	CALL				; std.Print(42)
	RETURN

zero:	NOP				; () -> i32
	PUSH_I32 0
	RETURN

life:	NOP				; () -> i32
	PUSH_FN std.Add_I32
	PUSH_I32 2
	PUSH_I32 40
	FEED 2				; std.Add_I32.Feed(2, 40)
	CALL				; std.Add_I32(2, 40)
	RETURN