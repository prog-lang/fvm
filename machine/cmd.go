package machine

import "fmt"

// Cmd (a.k.a. "Command") is a
//
//  1. sequence of executable instructions
//  2. that follows the Function interface
//  3. and can be treated like a regular Object
//
// It is used to contain compiled bytecode instructions generated from the
// human-written source code.
type Cmd struct {
	data  Data
	code  Code
	stack *Stack[Object]
	args  []Object
	ip    uint64
	argc  uint64
	done  bool
}

type Data interface {
	ReadAt(addr, length uint64) []uint8
}

type Code interface {
	Fetch(addr uint64) Do
	FetchInstruction(addr uint64) (opcode uint64, operand []uint8)
}

type Do func(*Cmd)

func MakeCmd(data Data, code Code, ip, argc uint64) Cmd {
	return Cmd{
		data: data,
		code: code,
		ip:   ip,
		argc: argc,
		args: make([]Object, 0),
	}
}

func (cmd Cmd) Feed(arg Object) Object {
	cmd.args = append(cmd.args, arg)
	if uint64(len(cmd.args)) >= cmd.argc {
		return cmd.call()
	}
	return cmd
}

func (cmd Cmd) call() Object {
	cmd.stack = NewStack[Object](Unit{})
	//*                         ^^^^^^^^
	//? Unit{} is the default return value of any Cmd.
	//? It is pushed to the bottom of the data stack by default in order to
	//? forego return emptiness checks.

	for !cmd.done {
		cmd.code.Fetch(cmd.ip)(&cmd)
		cmd.ip += SizeInstruction
	}

	return cmd.stack.Pop()
}

func (cmd Cmd) String() string {
	return fmt.Sprintf("Cmd<%d>(%d/%d)", cmd.ip, len(cmd.args), cmd.argc)
}
