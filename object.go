package machine

type Object any

type Unit struct{}

type Function interface {
	Object
	Feed(Object) Function
	Call() Object
}
