package machine

import (
	"fmt"
	"strconv"

	"github.com/prog-lang/pure/std"

	"github.com/charmbracelet/log"
)

var stdlib = [std.Count]Fn{
	id,
	iff,

	/* I32 operations */

	add_I32,
	sub_I32,
	mul_I32,
	div_I32,
	show_I32,

	/* Cmd operations */

	cmd_Cmd,
	// async_Cmd,
	map_Cmd,
	swap_Cmd,
	then_Cmd,
	chain_Cmd,

	/* I/O operations */

	print_Str,
}

var (
	id = MakeFn(
		1,
		func(args []Object) Object {
			debug("id", args)
			return args[0]
		},
	)
	iff = MakeFn(
		3,
		func(args []Object) Object {
			debug("if", args)
			condition := args[0].(bool)
			first := args[1]
			second := args[2]
			if condition {
				return first
			} else {
				return second
			}
		},
	)

	/* I32 operations */

	add_I32 = MakeFn(
		2,
		func(args []Object) Object {
			debug("add[i32]", args)
			return args[0].(int32) + args[1].(int32)
		},
	)
	sub_I32 = MakeFn(
		2,
		func(args []Object) Object {
			debug("sub[i32]", args)
			return args[0].(int32) - args[1].(int32)
		},
	)
	mul_I32 = MakeFn(
		2,
		func(args []Object) Object {
			debug("mul[i32]", args)
			return args[0].(int32) * args[1].(int32)
		},
	)
	div_I32 = MakeFn(
		2,
		func(args []Object) Object {
			debug("div[i32]", args)
			return args[0].(int32) / args[1].(int32)
		},
	)
	show_I32 = MakeFn(
		1,
		func(args []Object) Object {
			debug("show[i32]", args)
			return strconv.FormatInt(int64(args[0].(int32)), 10)
		},
	)

	/* Cmd operations */

	cmd_Cmd = MakeFn(
		1,
		func(args []Object) Object {
			debug("cmd", args)
			object := args[0]
			return Cmd(func() Object { return object })
		},
	)
	async_Cmd = MakeFn(
		1,
		func(args []Object) Object {
			debug("async", args)
			cmd := args[0].(Cmd)
			return Cmd(func() Object {
				go cmd.Exec()
				return Unit{}
			})
		},
	)
	map_Cmd = MakeFn(
		2,
		func(args []Object) Object {
			debug("map_Cmd", args)
			fn := args[0].(Function)
			cmd := args[1].(Cmd)
			return Cmd(func() Object {
				return fn.Apply(cmd.Exec())
			})
		},
	)
	swap_Cmd = MakeFn(
		2,
		func(args []Object) Object {
			debug("swap_Cmd", args)
			object := args[0]
			cmd := args[1].(Cmd)
			return Cmd(func() Object {
				_ = cmd.Exec()
				return object
			})
		},
	)
	then_Cmd = MakeFn(
		2,
		func(args []Object) Object {
			debug("then", args)
			cmd0 := args[0].(Cmd)
			cmd1 := args[1].(Cmd)
			return Cmd(func() Object {
				_ = cmd0.Exec()
				return cmd1.Exec()
			})
		},
	)
	chain_Cmd = MakeFn(
		2,
		func(args []Object) Object {
			debug("chain", args)
			cmd := args[0].(Cmd)
			fn := args[1].(Function)
			return Cmd(func() Object {
				return fn.Apply(cmd.Exec())
			})
		},
	)

	/* I/O operations */

	print_Str = MakeFn(
		1,
		func(args []Object) Object {
			debug("print", args)
			return Cmd(func() Object {
				str := args[0].(string)
				fmt.Print(str)
				return str
			})
		},
	)
)

/* Utility functions */

func debug(name string, args []Object) {
	log.Debug(name, "args", args)
}
