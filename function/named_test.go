package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlankIdentifier(t *testing.T) {
	// multiple functions can be declared with names as the blank
	// identifier _, in which cases, the declared functions can never be called.
	assert.True(t, 1 == 1)
}

func _() {
}

func TestShadowReturnedNamedValue(t *testing.T) {
	v := func() (val int) {
		// no new variables on left side of :=
		// 说明，returned named value 并没有处在一个更大的scope之中
		// val := 10

		if 1 == 1 {
			// the Go 1 compilers disallow(不允许) return statements without arguments
			// if any of the named return values is shadowed at the point of the return statement.
			// 所以，只写 return 是不行的
			val := 10
			return val
		}

		return
	}()
	assert.True(t, v == 10)
}
