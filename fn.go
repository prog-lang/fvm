package machine

type Fn struct {
	args []Object
	eval Eval
}

type Eval func(args []Object) Object

func MakeFn(eval Eval) Fn {
	return Fn{
		eval: eval,
	}
}

func (fn Fn) Feed(args []Object) Function {
	fn.args = append(fn.args, args...)
	return fn
}

func (fn Fn) Call() Object {
	return fn.eval(fn.args)
}
