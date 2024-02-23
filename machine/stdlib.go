package machine

import (
	"fmt"

	"github.com/prog-lang/pure/std"

	"github.com/charmbracelet/log"
)

var stdlib = [std.Count]Eval{
	func(args []Object) Object {
		log.Debug("id", "args", args)
		return args[0]
	},

	func(args []Object) Object {
		log.Debug("add[i32]", "args", args)
		return args[0].(int32) + args[1].(int32)
	},
	func(args []Object) Object {
		log.Debug("sub[i32]", "args", args)
		return args[0].(int32) - args[1].(int32)
	},
	func(args []Object) Object {
		log.Debug("mul[i32]", "args", args)
		return args[0].(int32) * args[1].(int32)
	},
	func(args []Object) Object {
		log.Debug("div[i32]", "args", args)
		return args[0].(int32) / args[1].(int32)
	},

	func(args []Object) Object {
		log.Debug("print", "args", args)
		fmt.Print(args...)
		return Unit{}
	},
}
