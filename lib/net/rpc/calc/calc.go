// Package rpc provides access to the exported methods of an object
// across a network or other I/O connection.
package calc

// the method's type is exported.
type Args struct {
	N, M int
}

// Only methods that satisfy these criteria will be made available for remote access
// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.
//
// func (t *T) MethodName(argType T1, replyType *T2) error
// where T1 and T2 can be marshaled by encoding/gob.
// NOTE: 输入和输出一定是可以被编码的，默认的编码方式gob
//       返回值error也会被返回，返回String()的值
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}

func (t *Args) Add(a int, ret *int) error {
	*ret = a + t.M + t.N
	return nil
}
