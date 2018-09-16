# golang-learning

不断更新自己对golang的认知

1. [variable](./variable): 介绍多重赋值、类型转换、**比较**、变量名与类型同名、作用域、零值和常量.
3. [control](./control): 介绍for、 if、 switch 和 type switch.
7. [error](./error): 介绍panic的工作原理，如何与recover结合使用；panic和error使用的边界；repanic
4. [func](./func): init函数以及初始化顺序, 可变参数、命名的返回值, closure, 高阶函数
11. [struct](./struct): 介绍如何通过匿名内部类实现继承，以及匿名类的其它用法
5. [import](./import)
6. [interface](./interface)
8. [chan](./chan)
12. [type](./type)
13. [routine](./routine)
14. [other](./other)
15. [标准库](./lib)

- Misusing pointers with value types

Passing a value as a parameter in a function or as receiver to a method may seem
a misuse of memory, because a value is always copied. 值传递看似浪费内存.

But on the other hand values are allocated on the stack, which is quick and relatively cheap. 值传递借助栈, 代价没那么大.

If you would pass a pointer to the value instead the Go compiler in most cases will see this as the making of an object,
and will move this object to the heap, so also causing an additional memory allocation. golang的内存管理原则

therefore nothing was gained in using a pointer instead of the value!

总结：分辨变量的生命周期是重要的，我们从一开始就应该决定是否应该将对象放在堆上。
对于临时变量或占用空间小的变量，可以采用值传递.
