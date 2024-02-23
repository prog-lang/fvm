package machine

type Cmd struct {
	data Data
	code Code
	ip   int32

	args []Object
	ds   *Stack[Object]
	done bool
}

type Data interface {
	ReadAt(addr, length int32) []uint8
}

type Code interface {
	Fetch(addr int32) Do
}

type Do func(*Cmd)

func MakeCmd(data Data, code Code, ip int32) Cmd {
	return Cmd{
		data: data,
		code: code,
		ip:   ip,
		args: make([]Object, 0),
	}
}

func (cmd Cmd) Feed(args []Object) Function {
	cmd.args = append(cmd.args, args...)
	return cmd
}

func (cmd Cmd) Call() Object {
	cmd.initStack()
	for !cmd.done {
		cmd.code.Fetch(cmd.ip)(&cmd)
		cmd.ip += InstructionSize
	}
	return cmd.returnValue()
}

func (cmd *Cmd) initStack() {
	cmd.ds = NewStack[Object](defaultDataStackCapacity).Push(Unit{})
	//*                                                ^^^^^^^^^^^^
	//? Unit{} is the default return value of any Cmd.
	//? It is pushed to the bottom of the data stack by default in order to
	//? forego return emptiness checks.
}

func (cmd *Cmd) returnValue() Object {
	return cmd.ds.Pop()
}
