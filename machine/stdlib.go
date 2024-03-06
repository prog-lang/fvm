package machine

import (
	"fmt"
	"strconv"

	"github.com/prog-lang/pure/std"

	"github.com/charmbracelet/log"
)

var stdlib = [std.Count]Fn{
	MakeFn(
		1,
		func(args []Object) Object {
			log.Debug("id", "args", args)
			return args[0]
		},
	),

	MakeFn(
		2,
		func(args []Object) Object {
			log.Debug("add[i32]", "args", args)
			return args[0].(int32) + args[1].(int32)
		},
	),
	MakeFn(
		2,
		func(args []Object) Object {
			log.Debug("sub[i32]", "args", args)
			return args[0].(int32) - args[1].(int32)
		},
	),
	MakeFn(
		2,
		func(args []Object) Object {
			log.Debug("mul[i32]", "args", args)
			return args[0].(int32) * args[1].(int32)
		},
	),
	MakeFn(
		2,
		func(args []Object) Object {
			log.Debug("div[i32]", "args", args)
			return args[0].(int32) / args[1].(int32)
		},
	),
	MakeFn(
		1,
		func(args []Object) Object {
			log.Debug("show[i32]", "args", args)
			return strconv.FormatInt(int64(args[0].(int32)), 10)
		},
	),

	MakeFn(
		1,
		func(args []Object) Object {
			log.Debug("print", "args", args)
			return Cmd(func() Object {
				str := args[0].(string)
				fmt.Print(str)
				return str
			})
		},
	),
}
