# File

File类型是一个与操作系统紧密关联的数据结构，我们无法弄清楚它的struct类型。与它紧密关联的两个类型FileMode和FileInfo。

> 注意，```func Pipe() (r *File, w *File, err error)``` 创建管道。

## FileMode

1. A FileMode represents a file's mode and permission bits. FileMode对应了Linux文件系统的文件类型和文件权限
2. The bits have the same definition on all systems, so that information about files can be moved from one system to another portably. 突出了os package的跨平台性
3. Not all bits apply to all systems.

## 设置常量的技巧

1. 不但是要定义常量；
2. 同时指定常量的类型，该类型有别于内置类型

```
type FileMode uint32
const (
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
    ModeAppend                                     // a: append-only
    ModeExclusive                                  // l: exclusive use
    ModeTemporary                                  // T: temporary file (not backed up)
    ModeSymlink                                    // L: symbolic link
    ModeDevice                                     // D: device file
    ModeNamedPipe                                  // p: named pipe (FIFO)
    ModeSocket                                     // S: Unix domain socket
    ModeSetuid                                     // u: setuid
    ModeSetgid                                     // g: setgid
    ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
    ModeSticky                                     // t: sticky

    // Mask for the type bits. For regular files, none will be set.
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

    ModePerm FileMode = 0777 // Unix permission bits
)
```

> 不要直接操作FileMode，因为调用该类型的方法会更方便。

## FileInfo

```
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
```

## 打开文件操作

```
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error


func Create(name string) (*File, error)
func Open(name string) (*File, error)

const (
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)   // Flags to OpenFile
func OpenFile(name string, flag int, perm FileMode) (*File, error)

func Pipe() (r *File, w *File, err error)
```

### 为什么需要OpenFile函数

创建文件时，一般需要指定FileMode和permission。

打开文件时一般需要指定打开方式，包括：读、写或读写方式。

```
Create creates the named file with mode 0666 (before umask), truncating
it if it already exists. If successful, the associated file descriptor has mode O_RDWR.

Open opens the named file for reading. If successful, the associated file descriptor has mode O_RDONLY. 
```

可见，虽然Create和Open是最常用的两种打开方式，但是缺乏灵活性。FileMode，permission 和读写方式是固定写死的，而OpenFile的最大优势就在于它可以让调用者自己选择合适的组和方式。 注意，这也是FileMode的每一个标志占据一位的原因（方便组和）。

## 元数据操作

```
func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (f *File) Close() error
func (f *File) Fd() uintptr
func (f *File) Name() string
func (f *File) Stat() (FileInfo, error)
```

## 读文件操作

```
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) Sync() error
```

## 写文件操作

```
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)

func (f *File) Sync() error
    Sync commits the current contents of the file to stable storage.
    Typically, this means flushing the file system's in-memory copy of
    recently written data to disk.

```
