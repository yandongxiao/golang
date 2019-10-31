a function prototype: func Double(n int) (result int)
a function declaration: func Double(n int) (result int){}
a function signature: is composed of two type list, one is the input parameter type list, the other is the output result type lists.
a function type: is composed of the func keyword and a function signature literal.

When we declare a custom function, we also declared an immutable function value acutally.

It is fatal error to call a nil function to start a new goroutine. The fatal error is not recoverable and will make the whole program crash.
For other situations, calls to nil function values will produce recoverable panics, including deferred function calls.
