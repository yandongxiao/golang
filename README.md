# golang-learning

不断更新自己对golang的认知

1. [variable](./variable): 介绍多重赋值、类型转换、**比较**、变量名与类型同名、作用域、零值和常量.
2. [control](./control): 介绍for、 if、 switch 和 type switch.
3. [error](./error): 介绍panic的工作原理，如何与recover结合使用；panic和error使用的边界；repanic.
4. [func](./func): 介绍init函数以及初始化顺序, 可变参数、命名的返回值, closure, 高阶函数
5. [struct](./struct): 介绍继承、封装、匿名内部类等属性.
6. [type](./type): 介绍map, slice, array, string等数据结构
7. [interface](./interface): 介绍interface的工作原理，比较方法，多态，type assertion等.
8. [chan](./chan): 介绍chan的各方面属性，如何使用routine+chan实现协程间同步，以及常见的同步手段.
13. [routine](./routine):
15. [标准库](./lib): 介绍golang各个模块的使用方法, 例如sync, io, os,strings等

- Misusing pointers with value types

Passing a value as a parameter in a function or as receiver to a method may seem
a misuse of memory, because a value is always copied. 值传递看似浪费内存.

But on the other hand values are allocated on the stack, which is quick and relatively cheap. 值传递借助栈, 代价没那么大.

If you would pass a pointer to the value instead the Go compiler in most cases will see this as the making of an object,
and will move this object to the heap, so also causing an additional memory allocation. golang的内存管理原则

therefore nothing was gained in using a pointer instead of the value!

总结：分辨变量的生命周期是重要的，我们从一开始就应该决定是否应该将对象放在堆上。
对于临时变量或占用空间小的变量，可以采用值传递.
