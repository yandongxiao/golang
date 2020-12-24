# Go 1.1 is released

The most significant improvements are performance-related.
We have made optimizations in the compiler and linker,
garbage collector, goroutine scheduler, map implementation,
and parts of the standard library.

It is likely that your Go code will run noticeably faster when built with Go 1.1.

## 摘要

1. add method value, like a.Foo. And method expresstion like A.Foo
2. the addition of a race detector
3. expand stand library, changes to go command
4. remains compatible with Go 1.0. 保持了完全兼容
5. Both the gc and gccgo implementations now make int and uint 64 bits on 64-bit platforms

## Unicode

To make it possible to represent code points greater than 65535 in UTF-16, Unicode defines surrogate halves,
a range of code points to be used only in the assembly of large values, and only in UTF-16.
Unicode为了UTF-16编码，竟然在Unicode当中定义了一些surrogate halves.

The code points in that surrogate range are illegal for any other purpose.

In Go 1.1, this constraint is honored by the compiler, libraries, and run-time.
A surrogate half is illegal as a rune value, when encoded as UTF-8. 相当于Unicode转UTF-8.
It is treated as an encoding error and will yield the replacement rune, utf8.RuneError, U+FFFD.

处理方式：

1. constants such as '\ud800' and "\ud800" are now rejected by the compilers.
2. "\xed\xa0\x80"的形式导致编译器无法识别，但是解码未Unicode时，仍返回U+FFFD.

## Performance

1. The gc compilers generate better code in many cases, most noticeably for floating point on the 32-bit Intel architecture.
2. The gc compilers do more in-lining, including for some operations in the run-time such as append and interface conversions.
3. There is a new implementation of Go maps with significant reduction in memory footprint and CPU time
4. The garbage collector has been made more parallel, which can reduce latencies for programs running on multiple CPUs.
5. The garbage collector is also more precise, which costs a small amount of CPU time but can reduce the size of the heap significantly, especially on 32-bit architectures.
6. Due to tighter coupling of the run-time and network libraries, fewer context switches are required on network operations.


# Go 1.2 is released

Go 1.2 includes a couple of **minor** language changes, several
improvements to the language implementation and tools, some performance improvements,
and many additions and (backward-compatible) changes to the standard library.

## 摘要

1. 通过二进制或源代码安装
2. A new three-index slice syntax adds the ability to specify capacity as well as length. ```slice = array[2:4:7]```
   This allows the programmer to pass a slice value that can only access a limited portion of the underlying array.
3. compute and display test **coverage** results
4. Goroutines are now pre-emptively **scheduled**
5. Limit on the number of threads
6. Gopher Academy
7. keeps the promise of compatibility

## 协程优化

### Pre-emption in the scheduler

为什么说是部分解决：如果for{}内部没有非内敛函数调用，那么该goroutine会一直霸占着这个线程

In prior releases, a goroutine that was looping forever could starve out other goroutines on the same thread,
a serious problem when GOMAXPROCS provided only one user thread.
In Go 1.2, this is partially addressed: **The scheduler is invoked occasionally upon entry to a function**.
This means that any loop that includes a (non-inlined) function call can be pre-empted, allowing other goroutines to run on the same thread.

### Limit on the number of threads

goroutines are multiplexed onto threads
Go 1.2 introduces a configurable limit (default 10,000) to the total number of threads

### Stack size

In Go 1.2, the minimum size of the stack when a goroutine is created has been lifted from 4KB to 8KB. 这不是一个最终的完美的解决方案

At the other end, the new function SetMaxStack in the runtime/debug package controls the maximum size of a single goroutine's stack.

The default is 1GB on 64-bit systems and 250MB on 32-bit systems.


# Go 1.3 is released

This release comes six months after our last major release and
provides better performance, improved tools,
support for running Go in new environments, and more.

1. Static analysis features of godoc
2. 支持更多的操作系统，虽然还处在实验阶段，如DragonFly BSD, Plan 9, and Solaris
3. Changes to the runtime have improved the performance
4. The garbage collector is now precise when examining stacks
5. Allocate each Go routine a contiguous piece of memory for its stack, grown by reallocation/copy when it fills up.
   "segmented" model to a contiguous model.
6. For a while now(Go 1.0), the garbage collector has been precise when examining values in the heap;
   the Go 1.3 release adds equivalent precision to values on the stack.


## stack

- For a while now(Go 1.0), the garbage collector has been precise when examining values in the heap
- the Go 1.3 release adds equivalent precision to values on the **stack**.
- This means that a non-pointer Go value such as an integer will never be mistaken for a pointer and prevent unused memory from being reclaimed.
- Starting with Go 1.3, the runtime assumes that values with pointer type contain pointers and other values do not. This assumption is fundamental to the precise behavior of both stack expansion and garbage collection. 
- Programs that use package unsafe to store integers in pointer-typed values are illegal and will crash if the runtime detects the behavior.
- Programs that use package unsafe to store pointers in integer-typed values are also illegal but more difficult to diagnose during execution.
- Because the pointers are hidden from the runtime, a stack expansion or garbage collection may reclaim the memory they point at, creating dangling pointers.


# Go 1.4 is released

It contains a small language change, support for more operating systems and processor architectures, and improvements to the tool chain and libraries.
As always, Go 1.4 keeps the promise of compatibility, and almost everything will continue to compile and run without change when moved to 1.4.

## 摘要

1. The most notable new feature in this release is official support for **Android**.
2. go generate: automate the running of tools to generate source code before compilation.
3. Improving the garbage collector and preparing the ground for a fully concurrent collector to be rolled out in the next few releases
4. much of the runtime code has been translated to Go
5. Other related changes also reduce the heap size, which is smaller by 10%-30% overall
6. interface value holds a pointer.
7. Canonical import paths(如何保证client使用的是官方Path)
   To complement this new feature, a check has been added at update time to verify that the local package's remote repository
   matches that of its custom import. The go get -u command will fail to update a package if its remote repository has changed
   since it was first downloaded. The new -f flag overrides this check.

## runtime

- GC

Prior to Go 1.4, the runtime (garbage collector, concurrency support, interface management, maps, slices, strings, ...) was mostly written in C, with some assembler support.

In 1.4, much of the code has been translated to Go so that the garbage collector can scan the stacks of programs in the runtime
and get accurate information about what variables are active.

This rewrite allows the garbage collector in 1.4 to be fully precise, meaning that it is aware of the location of all active pointers in the program.
This means the heap will be smaller as there will be no false positives keeping non-pointers alive.
Other related changes also reduce the heap size, which is smaller by 10%-30% overall relative to the previous release.

- stack

Performance can be noticeably better in some cases and is always more predictable.

the default starting size for a goroutine's stack in 1.4 has been reduced from 8192 bytes to 2048 bytes.

- interface value

The implementation of interface values has been modified.
In earlier releases, the interface contained a word that was either a pointer or a one-word scalar value, depending on the type of the concrete object stored.
This implementation was problematical for the garbage collector, so as of 1.4 interface values always hold a pointer.

In running programs, most interface values were pointers anyway, so the effect is minimal, but programs that store integers (for example) in interfaces will see more allocations.

- performance

The garbage collector was sped up, leading to measurable improvements for garbage-heavy programs.

On the other hand, the new write barriers slow things down again, typically by about the same amount but, depending on their behavior,
some programs may be somewhat slower or faster.


# Go 1.5 is released

- This release includes **significant** changes to the implementation.
- 19 August 2015

## 摘要

### 重大改变

1. The compiler and runtime are now written entirely in Go (with a little assembler).
   C is no longer involved in the implementation, and so the C compiler that was once necessary for building the distribution is gone.
   在1.4版本中，大部分run time相关的代码由go语言重构（e.g. 堆栈管理器，垃圾回收管理器，并发管理等）
2. The garbage collector was **completely redesigned**. The garbage collector is now concurrent and provides dramatically
   lower pause times by running, when possible, in parallel with other goroutines.
   The "stop the world" phase of the collector will almost always be under 10 milliseconds and usually much less(less than 2ms).
3. Related improvements to the scheduler allowed us to change the default GOMAXPROCS value. in prior releases it defaulted to 1, 现在是CPU的核数
4. Changes to the linker enable building Go packages into archives or shared libraries that may be linked into or loaded by C programs

### 其它重要改变

1. Support for "internal" packages. 避免内部的package被其它项目使用
2. a standard mechanism for managing dependencies in Go programs. 尚处于实验阶段
3. The new "go tool trace" command enables the visualisation of program traces.
4. The new "go doc" command provides an improved command-line interface for viewing Go package documentation.
5. language change: the lifting of a restriction in the map literal syntax.

## Go 1.5 Release Notes

### the lifting of a restriction in the map literal syntax

Due to an oversight, and now

```
m := map[Point]string{
    {29.935523, 52.891566}:   "Persepolis",
    {-25.352594, 131.034361}: "Uluru",
    {37.422455, -122.084306}: "Googleplex",
}
```

### Why the compiler and runtime are now written entirely in Go

- "gc" Go toolchain如下:
    - C program to parse Go code to C/Assembly langurage --> Plan 9 compiler toolchain.
    - assemblers, C compilers, and linkers are adopted essentially unchanged

- Why use 'gc' compiler at start
    - Go did not exist
    - once Go did exist, it often changed in significant
    - 以上问题解决以后，为了Go编译器的长远发展，决定重写。
    - applying an automatic translator
    - go tool compile
    - go too linke
    - go tool asm
    - 源代码安装的时候， a Go compiler must be available to compile the distribution from source.

### Garbage collector

Details of the new collector were presented in a talk at GopherCon 2015.

### Runtime

- In Go 1.5, the order in which goroutines are scheduled has been changed.
- The properties of the scheduler were never defined by the language.
- sets the default number of **threads** to run simultaneously, defined by GOMAXPROCS, to the number of cores available on the CPU.
- The default setting of GOMAXPROCS in all extant Go releases is 1, because programs with frequent goroutine switches ran much slower when using multiple threads.
- It is much cheaper to switch between two goroutines in the same thread than to switch between two goroutines in different threads.
- Goroutine **scheduling affinity** and other improvements to the scheduler have largely addressed the problem


# Go 1.6 is released

Although the release of Go 1.5 six months ago contained dramatic implementation changes, this release is more incremental.

- The most significant change is support for HTTP/2 in the net/http package.
    - In Go 1.6, support for HTTP/2 is enabled by default for both servers and clients when using HTTPS
- The runtime has added lightweight, best-effort detection of concurrent misuse of maps.
    - If one goroutine is writing to a map, no other goroutine should be reading or writing the map concurrently.
    - If the runtime detects this condition, it prints a diagnosis and crashes the program.(这个狠)
- Garbage-collection pauses are even lower than with Go 1.5, but this is particularly noticeable for programs using large amounts of memory.
- Go 1.5 introduced experimental support for a “vendor” directory that was enabled by an environment variable. In Go 1.6, the feature is now enabled by default.

## cgo

The major change is the definition of rules for sharing Go pointers with C code, to ensure that such C code can coexist with Go's garbage collector.


# Go 1.7 is released

There are several significant changes in this release

- compiler improvements
- the addition of the context package
- support for hierarchical tests and benchmarks
- 170 people contributed to this release, including 140 from the Go community.

## compiler improvements

- A new compiler back end, based on static single-assignment form (SSA), has been under development for the past year.
- By representing a program in SSA form, a compiler may perform advanced optimizations more easily.
- This new back end generates more compact, more efficient code that includes optimizations like bounds check elimination and common subexpression elimination.
- users have observed a significant speedup in compile time and a reduction in binary size by as much as 20–30%.
- We observed a 5–35% speedup across our benchmarks[https://golang.org/test/bench/go1/] 编译、二进制、运行时间均有提升
- Programs should run a bit faster due to speedups in the garbage collector and optimizations in the standard library.
- Programs with many idle goroutines will experience much shorter garbage collection pauses than in Go 1.6.

## Work done in Go 1.7

- The first is the new SSA backend that was enabled for AMD64 in this release. While the primary motivation for SSA was improved performance, the better generated code is also smaller.
    - The SSA backend shrinks Go binaries by ~5%

- The second change is method pruning.
    - Now the compiler discards any unexported methods that do not match an interface.
    - Similarly the linker can discard other exported methods, those that are only accessible through reflection, if the corresponding reflection features are not used anywhere in the program.
    - That change shrinks binaries by 5–20%.

- The SSA backend shrinks Go binaries by ~5%
    - The new format shrinks Go binaries by a further 5–15%

## sub-tests and sub-benchmarks

### 关于并行

- Each test is associated with a test function.
- A test is called a parallel test if its test function calls the Parallel method on its instance of testing.T.
- A parallel test never runs concurrently with a sequential test. 并行测试永远不会与顺序测试并发执行
- its execution is suspended until its calling test function, that of the parent test, has returned.
- The -parallel flag defines the maximum number of parallel tests that can run in parallel.
- A test blocks until its test function returns and all of its subtests have completed. 
     - This means that the parallel tests that are run by a sequential test will complete before any other consecutive sequential test is run.
- NOTE: sub-test内部仍然可以调用Run，创建sub-sub-test

## HTTP Tracing

- a facility to gather fine-grained information throughout the lifecycle of an HTTP client request.
- Support for HTTP tracing is provided by the net/http/httptrace package.
- The httptrace package provides a number of hooks to gather information during an HTTP round trip about a variety of events.
    - Connection creation
    - Connection reuse
    - DNS lookups
    - Writing the request to the wire
    - Reading the response

HTTP tracing is a valuable addition to Go for those who are interested in debugging HTTP request latency and writing
tools for network debugging for outbound traffic. By enabling this new facility, we hope to see HTTP debugging, benchmarking
and visualization tools from the community — such as httpstat


# Go 1.8 is released

There are significant performance improvements and changes across the standard library.

- The compiler back end introduced in Go 1.7 for 64-bit x86 is now used on all architectures
- Compile times should be improved by about 15% over Go 1.7.
- Garbage collection pauses should be significantly shorter
    - usually under 100 microseconds and often as low as 10 microseconds. 通常情况下
    - 在内存量很大的情况下，也是低于1ms. 18Gbyte heap
- The HTTP server adds support for HTTP/2 Push, allowing servers to preemptively send responses to a client.
- It's now much simpler to sort slices using the newly added Slice function in the sort package.


## Changes to the language

- When explicitly converting a value from one struct type to another, as of Go 1.8 the tags are ignored.
- Thus two structs that differ **only in their tags** may be converted from one to the other.

```
func example() {
	type T1 struct {
		X int `json:"foo"`
	}
	type T2 struct {
		X int `json:"bar"`
	}
	var v1 T1
	var v2 T2
	v1 = T1(v2) // now legal
}
```


# Go 1.9 is released

## Abstract

- There are many changes to the language, standard library, runtime, and tooling.
- Most of the engineering effort put into this release went to improvements of the runtime and tooling. 工作重点
    - The most important change to the language is the introduction of type aliases. type T1 = T2
    - The sync package has added a new Map type, safe for concurrent access.
        - cache contention: when each core updates the count, it invalidates the local cache entries for that address in all the other ocres, and mark itself as the owner of the up-to-date value
        - the next core to update the count must fetch the value that the previous core wrote to its cache(CPU的缓存). 
        - that takes about 40ns.
        - 多个core同时写同一个地址，导致最后比一个core负责写还要慢, O(N)的时间复杂度.
        - It is not the best possible concurrency map for all use-cases.
            - stable keys
            - disjoint sotres. 即不同的core写入的是不同的key. 进一步，write once，read many times
            - concurrent loops. concurrent代表多协程写，loop表示每个协程都有很多的write操作.
            - 满足以上三种情况下，sync.Map才是最好的选择

    - The new Helper method, added to both testing.T and testing.B.
        - **Marks** the calling function as a test helper function
        - When the testing package prints file and line information, it shows the location of the call to a helper function instead of a line in the helper function itself.
    - The time package now **transparently** tracks monotonic time in each Time value.
        - 时间单调递增
        - elapsed := time.Since(start): this code now computes the right elapsed time even across a leap second clock reset
    - The Go compiler now supports compiling a package's functions in parallel
        - This is in addition to the go command's existing support for parallel compilation of separate packages. package之间的并行编译早已支持
    - GC
        - Library functions that used to trigger stop-the-world garbage collection now trigger concurrent garbage collection.
        -  runtime.GC, debug.SetGCPercent, and debug.FreeOSMemory, now trigger concurrent garbage collection, blocking only the calling goroutine until the garbage collection is done.
        - Large object allocation performance is significantly improved in applications using large (>50GB) heaps containing many large objects.
    - Moved GOROOT: GOROOT无需再提前设置

## Floating-point operators

```
// FMA allowed for computing r, because x*y is not explicitly rounded:
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)

// FMA disallowed for computing r, because it would omit rounding of x*y:
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z
```

总结：如果要保持精度，就直接使用x*y+z的方式.

```
To force the intermediate rounding, write float64(x*y) + z.
```

## Codebase Refactoring

### 代码重构的几个原因

1. The first reason is to split a package into more manageable pieces for users. 避免代码膨胀
2. The second reason is to improve naming. 抽象出更小的package，代码量越小，package内部命名越准确
3. The third reason is to lighten dependencies. 代码量越小，依赖的package就越少.
4. The fourth reason is to change the dependency graph so that one package can import another. 因为Go不允许有循环依赖

### 代码重构的步骤

1. First, introduce the new API. **interchangeable** with old API(上层应用最好能够无感知的迁移；两种数据类型可以互相操作)
2. Second, across as many commits as you need, convert all the uses of the old API to the new API.
3. Third, remove the old API.

相比atomic code repair, 分步骤重构的风险更低

在package之间移动常量，变量，函数都是可以的，都能保持interchangeable属性；唯独type不行。
因为type INT int的定义中，INT和int是两种不同的类型；它们尚可通过强制类型转换，保持interchangeable。但是struct是绝对不可能的

这也是type INT = int语法的由来


# Go 1.10 is released

1. improves caching of built packages, see go env GOCACHE
2. adds caching of successful test results
3. runs vet automatically during tests，静态分析，挖掘潜在的问题
4. permits passing string values directly between Go and C using cgo

# Go 1.11 is released

- two features stand out as being especially exciting: modules and WebAssembly support
    - modules: preliminary support, experimental phase
    - WebAssembly: experimental phase
        -  This allows programmers to compile Go programs to a binary format compatible with four major web browsers.
        - firefox, chrome, solaris, ie


- WebAssembly
    - 并不是一门编程语言，而是一份字节码标准，需要用高级编程语言编译出字节码放到 WebAssembly 虚拟机中才能运行(Go --> WebAssembly字节码 --> 主流浏览器支持WebAssembly VM)
    - WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine.
    - Wasm is designed as a portable target for compilation of high-level languages like C/C++/Rust, enabling deployment on the web for client and server applications.
    - 学习资料：https://github.com/golang/go/wiki/WebAssembly

- release notes
    - The runtime now uses a sparse heap layout so there is no longer a limit to the size of the Go heap (previously, the limit was 512GiB).