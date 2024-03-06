package machine

import "fmt"

// Proc (a.k.a. "Procedure") is a
//
//  1. sequence of executable instructions
//  2. that follows the Function interface
//  3. and can be treated like a regular Object
//
// It is used to contain compiled bytecode instructions generated from the
// human-written source code.
type Proc struct {
	data  Data
	code  Code
	ip    uint32
	argc  uint32
	args  []Object
	stack *Stack[Object]
	done  bool
}

type Data interface {
	ReadAt(addr, length uint32) []uint8
}

type Code interface {
	Fetch(addr uint32) Do
	FetchInstruction(addr uint32) (opcode uint32, operand []uint8)
}

type Do func(*Proc)

func MakeProc(data Data, code Code, ip, argc uint32) Proc {
	return Proc{
		data: data,
		code: code,
		ip:   ip,
		argc: argc,
		args: make([]Object, 0),
	}
}

func (proc Proc) Feed(arg Object) Object {
	proc.args = append(proc.args, arg)
	if uint32(len(proc.args)) >= proc.argc {
		return proc.call()
	}
	return proc
}

func (proc Proc) Exec() Object {
	return proc.Feed(Unit{})
}

func (proc Proc) call() Object {
	proc.stack = NewStack[Object](Unit{})
	//*                         ^^^^^^^^
	//? Unit{} is the default return value of any Proc.
	//? It is pushed to the bottom of the data stack by default in order to
	//? forego return emptiness checks.

	for !proc.done {
		proc.code.Fetch(proc.ip)(&proc)
		proc.ip += SizeInstruction
	}

	return proc.stack.Pop()
}

func (proc Proc) String() string {
	return fmt.Sprintf("Proc<%d>(%d/%d)", proc.ip, len(proc.args), proc.argc)
}
