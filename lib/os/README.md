# os

1. gives us a platform-independent interface to operating-system functionality;
2. its design is Unix-like;
3. it hides the differences between various operating systems to give a consistent view of files and other OS-objects.
4. The os interface is intended to be uniform across all operating systems. Features not generally available appear in the system-specific package syscall. 所以package os只是封装了操作系统最通用的部分
