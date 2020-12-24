// The delete built-in function deletes the element with the specified key
// (m[key]) from the map. If m is nil or there is no such element, delete
// is a no-op.
// NOTE: 只用于 map
package buildin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteExistedElement(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	delete(m, 1)

	// NOTE: the order is unpredictable even if
	// the same loop is run multiple times with
	// the same map
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 如果是Example的话需要这样来表示
	// Unordered output:
	// 2 2
	// 3 3

	assert.Equal(t, m, map[int]int{2: 2, 3: 3})
}

func TestNotExistedElement(t *testing.T) {
	m := map[int]int{1: 1, 2: 2}
	delete(m, 10)
	assert.Equal(t, m, map[int]int{1: 1, 2: 2})
}

func TestDeleteNilMap(t *testing.T) {
	var m map[int]int
	delete(m, 10)
	assert.Equal(t, m, map[int]int(nil))
	assert.NotEqual(t, m, map[int]int{})
}

// 删除当前正在遍历的元素，是安全的
// If map entries that have not yet been reached are removed during iteration,
// the corresponding iteration values will not be produced.
// If map entries are created during iteration, that entry may be produced during the iteration or may be skipped.
func TestDeleteMap(t *testing.T) {
	// NOTE: This is safe
	mm := map[string]int{"jack": 10, "bob": 20}
	for key := range mm {
		if key == "jack" {
			delete(mm, key)
		}
	}
	assert.Equal(t, mm, map[string]int{"bob": 20})
}
