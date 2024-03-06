package machine

import "fmt"

// Fn (a.k.a. "Builtin Function") is a special type that implement the Function
// interface. It is used to carry stdlib functions.
type Fn struct {
	argc uint32
	args []Object
	eval Eval
}

type Eval func(args []Object) Object

func MakeFn(argc uint32, eval Eval) Fn {
	return Fn{
		argc: argc,
		eval: eval,
	}
}

func (fn Fn) Apply(arg Object) Object {
	fn.args = append(fn.args, arg)
	if uint32(len(fn.args)) >= fn.argc {
		return fn.call()
	}
	return fn
}

func (fn Fn) call() Object {
	return fn.eval(fn.args)
}

func (fn Fn) String() string {
	return fmt.Sprintf("Fn(%d/%d)", len(fn.args), fn.argc)
}
