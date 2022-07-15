package zdpgo_type

import (
	"fmt"
	"testing"
)

// 判断是否全为true
func TestArrayBool_IsAllTrue(t *testing.T) {
	ab := NewArrayBool()
	ab.Append(true)
	ab.Append(true)
	ab.Append(true)
	fmt.Println(ab.IsAllTrue())
	ab.Append(false)
	fmt.Println(ab.IsAllTrue())
}
