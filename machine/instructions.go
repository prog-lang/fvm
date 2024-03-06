package machine

import (
	"github.com/prog-lang/pure/opcode"

	"github.com/charmbracelet/log"
)

// This array contains the entire instruction set of the Pure machine.
var instructions = [opcode.Count]func([]uint8) Do{
	func(_ []uint8) Do {
		return func(proc *Proc) {
			log.Debug("NOP")
		}
	},

	/* Stack manipulation */

	func(_ []uint8) Do {
		return func(proc *Proc) {
			log.Debug("PUSH_UNIT")
			proc.stack.Push(Unit{})
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			b := U8AsBool(operand[0])
			log.Debug("PUSH_BOOL", "bool", b)
			proc.stack.Push(b)
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			u8 := operand[0]
			log.Debug("PUSH_U8", "u8", u8)
			proc.stack.Push(u8)
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			i32 := U8x4AsI32(operand)
			log.Debug("PUSH_I32", "i32", i32)
			proc.stack.Push(i32)
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			addr := U8x4AsI32(operand)
			log.Debug("PUSH_FN", "@", addr)
			proc.stack.Push(stdlib[addr])
		}
	},
	func(operand []uint8) Do {
		// At IP we expect to see a NOP instruction. Its operand must specify
		// the argument count for the proc.
		//
		//     *-- opcode ---*-- operand --*
		// IP: | 00 00 00 00 | 03 00 00 00 |
		//     *-- NOP ------*-- argc(3) --*
		//
		return func(proc *Proc) {
			ip := U8x4AsU32(operand)
			_, operand := proc.code.FetchInstruction(ip)
			argc := U8x4AsU32(operand)
			log.Debug("PUSH_PROC", "@", ip, "argc", argc)
			proc.stack.Push(MakeProc(proc.data, proc.code, ip, argc))
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			index := U8x4AsU32(operand)
			log.Debug("PUSH_ARG", "#", index)
			proc.stack.Push(proc.args[index])
		}
	},
	func(operand []uint8) Do {
		return func(proc *Proc) {
			n := U8x4AsU32(operand)
			log.Debug("DROP", "n", n)
			proc.stack.Drop(n)
		}
	},

	/* Program flow */

	func(operand []uint8) Do {
		return func(proc *Proc) {
			argc := U8x4AsU32(operand)
			log.Debug("FEED", "argc", argc)
			args := proc.stack.Take(argc)
			object := proc.stack.Pop()
			for _, arg := range args {
				object = object.(Function).Feed(arg)
			}
			proc.stack.Push(object)
		}
	},
	func(_ []uint8) Do {
		return func(proc *Proc) {
			log.Debug("EXEC")
			proc.stack.Push(proc.stack.Pop().(Function).Feed(Unit{}))
		}
	},
	func(_ []uint8) Do {
		return func(proc *Proc) {
			log.Debug("RETURN")
			proc.done = true
		}
	},
}
