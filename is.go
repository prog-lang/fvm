package machine

import (
	"machine/opcode"

	"github.com/charmbracelet/log"
)

var IS = [opcode.Count]func([]byte) Do{
	func(operand []byte) Do {
		log.Debug("NOP")
		return func(cmd *Cmd) {}
	},

	/* Stack manipulation */

	func(operand []byte) Do {
		log.Debug("PUSH_UNIT")
		return func(cmd *Cmd) { cmd.ds.Push(Unit{}) }
	},
	func(operand []byte) Do {
		log.Debug("PUSH_BOOL")
		return func(cmd *Cmd) { cmd.ds.Push(U8AsBool(operand[0])) }
	},
	func(operand []byte) Do {
		log.Debug("PUSH_U8")
		return func(cmd *Cmd) { cmd.ds.Push(operand[0]) }
	},
	func(operand []byte) Do {
		log.Debug("PUSH_I32")
		return func(cmd *Cmd) { cmd.ds.Push(U8x4AsI32(operand)) }
	},
	func(operand []byte) Do {
		log.Debug("PUSH_CMD")
		return func(cmd *Cmd) {
			ip := U8x4AsI32(operand)
			cmd.ds.Push(NewCmd(cmd.data, cmd.code, ip))
		}
	},
	func(operand []byte) Do {
		log.Debug("DROP")
		return func(cmd *Cmd) { cmd.ds.Drop(int(U8x4AsI32(operand))) }
	},

	/* Program flow */

	func(operand []byte) Do {
		log.Debug("CALL")
		return func(cmd *Cmd) { cmd.ds.Push(cmd.ds.Pop().(Function).Call()) }
	},
	func(operand []byte) Do {
		log.Debug("BRANCH")
		return func(cmd *Cmd) {
			condition := cmd.ds.Pop().(bool)
			left := cmd.ds.Pop().(Function)
			right := cmd.ds.Pop().(Function)
			if condition {
				cmd.ds.Push(left)
			} else {
				cmd.ds.Push(right)
			}
		}
	},
	func(operand []byte) Do {
		log.Debug("RETURN")
		return func(cmd *Cmd) { cmd.done = true }
	},
}
