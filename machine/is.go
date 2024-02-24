package machine

import (
	"github.com/prog-lang/pure/opcode"

	"github.com/charmbracelet/log"
)

var is = [opcode.Count]func([]uint8) Do{
	func(_ []uint8) Do {
		log.Debug("NOP")
		return func(cmd *Cmd) {}
	},

	/* Stack manipulation */

	func(operand []uint8) Do {
		log.Debug("PUSH_UNIT")
		return func(cmd *Cmd) { cmd.stack.Push(Unit{}) }
	},
	func(operand []uint8) Do {
		b := U8AsBool(operand[0])
		log.Debug("PUSH_BOOL", "bool", b)
		return func(cmd *Cmd) { cmd.stack.Push(b) }
	},
	func(operand []uint8) Do {
		u8 := operand[0]
		log.Debug("PUSH_U8", "u8", u8)
		return func(cmd *Cmd) { cmd.stack.Push(u8) }
	},
	func(operand []uint8) Do {
		i32 := U8x4AsI32(operand)
		log.Debug("PUSH_I32", "i32", i32)
		return func(cmd *Cmd) { cmd.stack.Push(i32) }
	},
	func(operand []uint8) Do {
		addr := U8x4AsI32(operand)
		log.Debug("PUSH_FN", "@", addr)
		return func(cmd *Cmd) {
			eval := stdlib[addr]
			cmd.stack.Push(MakeFn(eval))
		}
	},
	func(operand []uint8) Do {
		ip := U8x4AsI32(operand)
		log.Debug("PUSH_CMD", "@", ip)
		return func(cmd *Cmd) {
			cmd.stack.Push(MakeCmd(cmd.data, cmd.code, ip))
		}
	},
	func(operand []uint8) Do {
		n := int(U8x4AsI32(operand))
		log.Debug("DROP", "n", n)
		return func(cmd *Cmd) { cmd.stack.Drop(n) }
	},

	/* Program flow */

	func(operand []uint8) Do {
		argc := U8x4AsI32(operand)
		log.Debug("FEED", "argc", argc)
		return func(cmd *Cmd) {
			args := cmd.stack.Take(int(argc))
			cmd.stack.Push(cmd.stack.Pop().(Function).Feed(args))
		}
	},
	func(_ []uint8) Do {
		log.Debug("CALL")
		return func(cmd *Cmd) { cmd.stack.Push(cmd.stack.Pop().(Function).Call()) }
	},
	func(_ []uint8) Do {
		log.Debug("BRANCH")
		return func(cmd *Cmd) {
			condition := cmd.stack.Pop().(bool)
			left := cmd.stack.Pop().(Function)
			right := cmd.stack.Pop().(Function)
			if condition {
				cmd.stack.Push(left)
			} else {
				cmd.stack.Push(right)
			}
		}
	},
	func(_ []uint8) Do {
		log.Debug("RETURN")
		return func(cmd *Cmd) { cmd.done = true }
	},
}
