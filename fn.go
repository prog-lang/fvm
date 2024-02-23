package machine

type Fn struct {
	args []Object
	eval Eval
}

type Eval func([]Object) Object

func MakeFn(eval Eval) Fn {
	return Fn{
		eval: eval,
	}
}

func (fn Fn) Feed(arg Object) Function {
	var args []Object
	copy(args, fn.args)
	fn.args = append(args, arg)
	return fn
}

func (fn Fn) Call() Object {
	return fn.eval(fn.args)
}
