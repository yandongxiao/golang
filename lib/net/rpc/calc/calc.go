// Package rpc provides access to the exported methods of an object across a network or other I/O connection.
// Only methods that satisfy these criteria will be made available for remote access
package calc

// the method's type is exported.
type Args struct {
	N, M int
}

// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}

func (t *Args) Add(a int, ret *int) error {
	*ret = a + t.M + t.N
	return nil
}
