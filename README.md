# golang-learning

把Go语言大致总结了一遍，大致包括：

    1. Go语言本身的特性
    2. 标准库
    3. Go的周边工具
    4. Go的Release Note信息

## How do I know whether a variable is allocated on the heap or the stack?

From a correctness standpoint, you don't need to know.
Each variable in Go exists as long as there are references to it.
The storage location chosen **by the implementation** is **irrelevant** to the semantics of the language.

The storage location does have an effect on writing efficient programs.
When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame.
However, if the compiler cannot **prove** that the variable is not referenced after the function returns,
then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.

Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap.
However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.
NOTE: new 也不一定保证就是在堆上

## Value Copy Costs

Generally speaking, the cost to copy a value is proportional to
the size of the value. However, value sizes are not the only factor
determining value copy costs. Different CPU architectures may
specially optimize value copying for values with specific sizes.

In practice, we can view values with sizes not larger than four
native words as small-size values. The costs of copying small-size
values are small. For the standard Go compiler, except values of
large-size struct and array types, most types in Go are small-size types.

One the other hand, we should also consider the fact that too many pointers
will increase the pressure of garbage collectors at run time.

We should also try to avoid using the two-iteration-variable forms
to iterate array and slice elements if the element types are large-size
types, for each element value will be copy to the second iteration
variable in the iteration process.
