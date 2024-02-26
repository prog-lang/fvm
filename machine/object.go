package machine

// Object is a special type that represents any value on which Pure machine can
// operate. For now, it is defined as "any" Go value.
type Object = any

// Unit is the basic most primitive Object used for convenience.
type Unit struct{}

type Function interface {
	// Pure machine regards Functions as Objects. Functions can, therefore, be
	// passed around as arguments or return values.
	Object

	// Functions consume arguments one-by-one. Feeding enough arguments into a
	// Function will result in its return value.
	//
	// Example:
	//
	//  add := stdlib[std.Add_I32] // add expects 2 arguments
	//  add1 := add.Feed(int32(1))
	//  result := add1.(Function).Feed(int32(2)) // result == int32(3)
	Feed(arg Object) Object
}
