package machine

import (
	"machine/opcode"

	"github.com/charmbracelet/log"
)

var IS = [opcode.Count]func([]byte) Do{
	func(operand []byte) Do {
		log.Debug("NOP")
		return func(fn *Routine) {}
	},

	/* Stack manipulation */

	func(operand []byte) Do {
		log.Debug("BOOL_PUSH")
		return func(fn *Routine) { fn.ds.Push(U8AsBool(operand[0])) }
	},
	func(operand []byte) Do {
		log.Debug("U8_PUSH")
		return func(fn *Routine) { fn.ds.Push(operand[0]) }
	},
	func(operand []byte) Do {
		log.Debug("I32_PUSH")
		return func(fn *Routine) { fn.ds.Push(U8x4AsI32(operand)) }
	},
	func(operand []byte) Do {
		log.Debug("DROP")
		return func(fn *Routine) { fn.ds.Drop(int(U8x4AsI32(operand))) }
	},

	/* Program flow */

	func(operand []byte) Do {
		log.Debug("JUMP")
		return func(fn *Routine) { fn.ip = U8x4AsI32(operand) }
	},
	func(operand []byte) Do {
		log.Debug("CALL")
		return func(fn *Routine) { fn.ds.Push(fn.ds.Pop().(Function).Call()) }
	},
	func(operand []byte) Do {
		log.Debug("BR")
		return func(fn *Routine) {
			condition := fn.ds.Pop().(bool)
			left := fn.ds.Pop().(Function)
			right := fn.ds.Pop().(Function)
			if condition {
				fn.ds.Push(left)
			} else {
				fn.ds.Push(right)
			}
		}
	},
	func(operand []byte) Do {
		log.Debug("DONE")
		return func(fn *Routine) { fn.done = true }
	},
}
