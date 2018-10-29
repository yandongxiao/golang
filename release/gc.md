# GC of golang

We manage the stacks and their size by copying them and updating pointers in the stack. It's a local operation so it scales fairly well.

 Go is a value-oriented language
 
 We also have a way ahead of time compilation system so the binary contains the entire runtime.