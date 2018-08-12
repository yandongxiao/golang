# os

1. gives us a platform-independent interface to operating-system functionality;
2. its design is Unix-like
3. Failing calls return values of type error rather than error numbers
4. it hides the differences between various operating systems to give a consistent view of files and other OS-objects
5. The os interface is intended to be uniform across all operating systems. Features not generally available appear in the system-specific package syscall. 所以package os只是封装了操作系统最通用的部分，借助package syscall来调用特殊系统的API

>
> package os提供的操作类似于一个个bash command一样简单易用。
> 

## PathError

```
type PathError struct {
    Op   string
    Path string
    Err  e
}

package os 自定义的错误类型，注意，它只是自定义了PathError.Err部分。
var (
    ErrInvalid    = errors.New("invalid argument") // methods on File will return this error when the receiver is nil
    ErrPermission = errors.New("permission denied")
    ErrExist      = errors.New("file already exists")
    ErrNotExist   = errors.New("file does not exist")
)

func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPathSeparator(c uint8) bool
func IsPermission(err error) bool
```

这种设计有两个好处：

1. 错误的比较只与文件是否存在有关系，与是哪个文件，如何操作文件无关，解耦。
2. 使得Is类函数更加通用，所有package os的错误类型，如```PathError```, ```LinkError```, ```SyscallError```都可以作为Is类函数的参数。

>
- 由于package os 定义了多中类型的错误，如PathError, LinkError等，每个函数的Documentation中都明确说明了返回的错误类型。