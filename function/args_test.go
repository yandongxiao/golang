package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: golang 不支持重载
func TestArgs(t *testing.T) {
	add := func(a int, b int) int {
		return a + b
	}
	assert.True(t, add(1, 2) == 3)

}

func TestArgVariadic(t *testing.T) {
	add := func(nums ...int) int { // nums 的类型为[]int
		sum := 0
		for _, x := range nums {
			sum += x
		}
		return sum
	}

	assert.True(t, add(1, 2, 3, 4) == 10)
	assert.True(t, add() == 0)
	x := []int{1, 2, 3, 4}
	assert.True(t, add(x...) == 10) // NOTE: x...的语法要求x必须是slice类型，不可以是数组类型
}
