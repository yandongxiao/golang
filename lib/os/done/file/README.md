# File

File类型是一个与操作系统紧密关联的数据结构，我们无法弄清楚它的struct类型。与它紧密关联的两个类型FileMode和FileInfo。

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

```
func Create(name string) (*File, error)
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func Pipe() (r *File, w *File, err error)

# 元数据
func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (f *File) Close() error
func (f *File) Fd() uintptr
func (f *File) Name() string
func (f *File) Stat() (FileInfo, error)

# 读
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) Sync() error

# 写
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)
```
