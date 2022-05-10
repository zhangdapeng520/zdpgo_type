package zdpgo_type

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/10 15:14
@Author : 张大鹏
@File : array_bool_test.go
@Software: Goland2021.3.1
@Description:
*/

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
