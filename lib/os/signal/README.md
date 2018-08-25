# signal

Linux信号是进程间通信的一种方式，基本流程如下：


- A进程发送指定信号给进程B

- 进程B一般会使用特定线程来处理信号

	1. 每一个线程都拥有一个signal mask，子进程继承的就是父进程的那个线程的signal mask。创建线程时，由```pthread_create```的参数指定signal mask，比如可以指定：The signal mask is inherited from the creating thread。
	2. 进程一启动就屏蔽所有信息（block all signal），这样主线程就不会接收到任何信号。新创建的线程默认也不会处理任何信号
	2. 创建指定的线程，对要处理的信号解除屏蔽（unblocked signal），指定信号处理函数。

- 当进程B接收到信号，会协调其它线程，妥善的完成相应的动作

## signal分类

- **Synchronous signals**

	Synchronous signals are signals triggered by errors in program execution: SIGBUS, SIGFPE, and SIGSEGV.

	These are only considered synchronous when caused by program execution, not when sent using os.Process.Kill or the kill program or some similar mechanism.

	In general, except as discussed below, **Go programs will convert a synchronous signal into a run-time panic**.

- **asynchronous signals**

    The remaining signals are asynchronous signals. They are not triggered
    by program errors, but are instead sent from the kernel or from some
    other program.

## 信号的默认处理方式

By default, a synchronous signal is converted into a run-time panic.

A SIGHUP, SIGINT(CTRL+C), or SIGTERM signal causes the program to exit. A SIGQUIT,
SIGILL, SIGTRAP, SIGABRT, SIGSTKFLT, SIGEMT, or SIGSYS signal causes the
program to exit with a stack dump.

A SIGTSTP, SIGTTIN, or SIGTTOU signal gets the system default behavior (these signals are used by the shell for job control).

The SIGPROF signal is handled directly by the Go runtime to implement runtime.CPUProfile.

Other signals will be caught but no action will be taken.


If the Go program is started with a non-empty signal mask, that will
generally be honored. However, some signals are explicitly unblocked:
the synchronous signals, SIGILL, SIGTRAP, SIGSTKFLT, SIGCHLD, SIGPROF,
and, on GNU/Linux, signals 32 (SIGCANCEL) and 33 (SIGSETXID) (SIGCANCEL
and SIGSETXID are used internally by glibc). **所以，signal mask不是完全继承自父进程的**

Subprocesses started by os.Exec, or by the os/exec package, will inherit the modified signal mask.

## 信号处理函数

Notify disables the default behavior for a given set of asynchronous signals and instead delivers them over **one or more** registered channels.

It is allowed to call Notify multiple times with different channels and the same signals: each channel receives copies of incoming signals independently.

