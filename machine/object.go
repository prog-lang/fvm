package machine

type Object = any

type Unit struct{}

type Function interface {
	Object
	Feed(args []Object) Function
	Call() Object
}
