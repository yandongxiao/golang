# [Go 1.3+ Compiler Overhaul](https://docs.google.com/document/d/1P3BLR31VA8cvLJLfMibSuTdwTuF7WWLux71CYD0eeD8/edit)

## Abstract

当前，Go编译器是用C语言编写的应用程序，它的主要工作是将Go语言转换成汇编代码.

Go的常见编译器有两个: gc 和 gccgo. gc编译器是继承自Plan 9，包括汇编器、C编译器、连接器。

所以：构建Go程序的过程: go程序 --> 中间结果(C语言) --> 汇编 --> 编译/链接. gccgo/gc负责将go代码装换为C代码

## background

1. The “gc” Go toolchain is derived from the Plan 9 compiler toolchain.
2.  The assemblers, C compilers, and linkers are adopted essentially unchanged
3. the Go compilers (in cmd/gc, cmd/5g, cmd/6g, and cmd/8g) are new C programs that fit into the toolchain.

- 为什么一开始使用C语言编写Go编译器

1. 我们在创建Go语言的过程中，只能使用C语言来开发编译器
2. Go语言创建至今，经常会有较大的特性调整. 可能会出现bootstrapping problem.

- 为什么使用Go语言来重新构建Go编译器

1. It is easier to write correct Go code than to write correct C code.
2. It is easier to debug incorrect Go code than to debug incorrect C code.
3. Work on a Go compiler necessarily requires a good understanding of Go.
   Implementing the compiler in C adds an unnecessary second requirement.
4. Go makes parallel execution trivial compared to C.
5. Go has better standard support than C for modularity, for automated rewriting, for unit testing, and for profiling.
6. Go is much more fun to use than C.

## Proposed Plan

通过开发一个translator，将Go编译器gc，由C语言转化为Go语言。这项工作由以下几个阶段组成：

1. Develop and debug the translator.
   写一个通用的C-->Go的翻译器是困难的，需要耗费大量的工作。毕竟，有些C语言特性，如macros, unions, and bit fields，在Go语言中无法找到对应的原语。
   但是如果把目标限制在对编译器的翻译，还是比较可行的.
2. 执行编译器替换。将C代码副本删除，以Go代码替换。这时的Go编译器是一个很“C语言”的Go应用程序。
3. 将翻译的编译器代码进行调整，如模块化，添加测试等，将Go编译器代码编程一个native Go应用程序
4. 进一步的代码优化
