package machine

type Routine struct {
	data Data
	code Code
	ip   int32

	args []Object
	ds   *Stack[Object]
	done bool
}

type Data interface {
	ReadAt(addr, length int32) []byte
}

type Code interface {
	Fetch(addr int32) Do
}

type Do func(*Routine)

func NewRoutine(data Data, code Code, ip int32) Routine {
	return Routine{
		data: data,
		code: code,
		ip:   ip,
		args: make([]Object, 0),
	}
}

func (fn Routine) Feed(arg Object) Function {
	var args []Object
	copy(args, fn.args)
	fn.args = append(args, arg)
	return fn
}

func (fn Routine) Call() Object {
	fn.ds = NewStack[Object](defaultDataStackCapacity)
	for !fn.done {
		fn.code.Fetch(fn.ip)(&fn)
		fn.ip += InstructionSize
	}
	return fn.ret()
}

func (fn *Routine) ret() Object {
	if fn.ds.Empty() {
		return Unit{}
	}
	return fn.ds.Pop()
}
