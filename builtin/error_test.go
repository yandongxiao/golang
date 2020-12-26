// errors are the last return value and have type error,
// a built-in interface. A nil value in the error position
// indicates that there was no error.
//
// package os 的 PathError 包含 Op, Path, Err.
// 构成的 error string: open /etc/passwx: no such file or directory.
// 调用 os.Open 的函数一般要么打印error string，要么向上返回该error.
// 所以，识别error抛出的源头是重要的.
//
// error strings 的格式：should identify their origin, such as
// by having a prefix naming the operation or package that
// generated the error.
//
// Error类型的命名规则: Error types **end** in "Error"
// and error variables are called (or **start** with) "err" or "Err".
package buildin

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsNew(t *testing.T) {
	err := errors.New("gerror: helloworld")
	assert.Equal(t, err.Error(), "gerror: helloworld")
}

func TestFmtErrorf(t *testing.T) {
	err := fmt.Errorf("gerror: %v", "helloworld")
	assert.Equal(t, err.Error(), "gerror: helloworld")
}

func TestFmtErrorf_W(t *testing.T) {
	err1 := fmt.Errorf("err1: %v", "helloworld")
	err2 := fmt.Errorf("err2: %w", err1)
	assert.Equal(t, err2.Error(), "err2: err1: helloworld")
	assert.True(t, errors.Is(err2, err1))
}

// A package can also define its own specific Error
// with additional methods, like net.Error.
type argError struct {
	arg  int
	prob string
}

func (err *argError) Error() string {
	return err.prob // NOTE: 返回的信息应该尽量全
}

func TestCustomizeError(t *testing.T) {
	var err error
	err = &argError{400, "helloworld"}

	// 如何防止type assertion出错时导致程序崩溃
	ae1, ok := err.(*argError)
	assert.True(t, ok)
	assert.Equal(t, err.Error(), "helloworld")

	err = fmt.Errorf("gerror: %w", err)
	assert.Equal(t, err.Error(), "gerror: helloworld")

	var ae2 *argError
	assert.True(t, errors.As(err, &ae2))
	assert.Equal(t, ae1, ae2)
}
