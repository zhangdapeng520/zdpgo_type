package zdpgo_type

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/4 21:41
@Author : 张大鹏
@File : set_test
@Software: Goland2021.3.1
@Description: 测试集合类型
*/

func TestSet_basic(t *testing.T) {
	s := NewSet()
	s.Add(1, 2, 3, 3, 2, 1)
	fmt.Println(s.ToSlice())

	s.Remove(2, 3)
	fmt.Println(s.ToSlice())
}
