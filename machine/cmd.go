package machine

// Cmd (a.k.a. "Command") is a special kind of a Function that represents
// some sort of communicational logic (as opposed to computational logic).
//
// Given function f defined as
//
//	f :: I32 -> I32;
//	f := x -> x + 2;
//
// f is a (pure) computational function. Its "reach" is limited to its own
// arguments. Whereas
//
//	http.get :: Str -> Cmd http.Response;
//
// http.get is an (impure) communicational function that relies on the internet
// connection.
//
// Other languages might type our http.get function as
//
//	http.get :: Str -> http.Response;
//
// but Pure is a purely functional programming language! Therefore, our promise
// to the programmer is that any Pure function, given the same arguments, is
// going to always return the same value. If we were to say that http.get simply
// returns an http.Response, we would break that promise, since, http.get can
// return a different http.Response for the same URL (as it depends on the logic
// of the server on the other side).
//
// As a rule of thumb, any function that performs I/O of any kind is going to
// have to return a Cmd. It makes sense - the http.Response might be different
// but the logic behind retrieving it is the same. Cmd is the type that contains
// said logic within itself!
type Cmd func() Object

func (io Cmd) Apply(arg Object) Object {
	return io()
}

func (io Cmd) Exec() Object {
	return io()
}
