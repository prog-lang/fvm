package machine

import (
	"github.com/prog-lang/pure/opcode"

	"github.com/charmbracelet/log"
)

// This array contains the entire instruction set of the Pure machine.
var instructions = [opcode.Count]func([]uint8) Do{
	func(_ []uint8) Do {
		return func(cmd *Cmd) {
			log.Debug("NOP")
		}
	},

	/* Stack manipulation */

	func(_ []uint8) Do {
		return func(cmd *Cmd) {
			log.Debug("PUSH_UNIT")
			cmd.stack.Push(Unit{})
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			b := U8AsBool(operand[0])
			log.Debug("PUSH_BOOL", "bool", b)
			cmd.stack.Push(b)
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			u8 := operand[0]
			log.Debug("PUSH_U8", "u8", u8)
			cmd.stack.Push(u8)
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			i32 := U8x4AsI32(operand)
			log.Debug("PUSH_I32", "i32", i32)
			cmd.stack.Push(i32)
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			addr := U8x4AsI32(operand)
			log.Debug("PUSH_FN", "@", addr)
			cmd.stack.Push(stdlib[addr])
		}
	},
	func(operand []uint8) Do {
		// At IP we expect to see a NOP instruction. Its operand must specify
		// the argument count for the Cmd.
		//
		//     *-- opcode ---*-- operand --*
		// IP: | 00 00 00 00 | 03 00 00 00 |
		//     *-- NOP ------*-- argc(3) --*
		//
		return func(cmd *Cmd) {
			ip := U8x4AsU32(operand)
			_, operand := cmd.code.FetchInstruction(ip)
			argc := U8x4AsU32(operand)
			log.Debug("PUSH_CMD", "@", ip, "argc", argc)
			cmd.stack.Push(MakeCmd(cmd.data, cmd.code, ip, argc))
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			index := U8x4AsU32(operand)
			log.Debug("PUSH_ARG", "#", index)
			cmd.stack.Push(cmd.args[index])
		}
	},
	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			n := U8x4AsU32(operand)
			log.Debug("DROP", "n", n)
			cmd.stack.Drop(n)
		}
	},

	/* Program flow */

	func(operand []uint8) Do {
		return func(cmd *Cmd) {
			argc := U8x4AsU32(operand)
			log.Debug("FEED", "argc", argc)
			args := cmd.stack.Take(argc)
			object := cmd.stack.Pop()
			for _, arg := range args {
				object = object.(Function).Feed(arg)
			}
			cmd.stack.Push(object)
		}
	},
	func(_ []uint8) Do {
		return func(cmd *Cmd) {
			log.Debug("BRANCH")
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
		return func(cmd *Cmd) {
			log.Debug("RETURN")
			cmd.done = true
		}
	},
}
