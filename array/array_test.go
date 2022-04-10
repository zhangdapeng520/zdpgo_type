package array

import (
	"fmt"
	"testing"
)

// 测试数组的基本使用
func TestArray_basic(t *testing.T) {
	arr := NewArray(1, 2, 3, 4)
	fmt.Println(arr.String())

	// 根据索引读取元素
	fmt.Println(arr.Get(0), arr.Get(1))

	// 更新元素
	arr.Set(0, 33)
	fmt.Println(arr.String())

	// 尾部插入
	arr.Append(111)
	fmt.Println(arr.String())

	// 中间插入
	arr.Insert(1, 44)
	fmt.Println(arr.String())

	// 头部插入
	arr.Insert(0, 100)
	fmt.Println(arr.String())

	// 删除元素
	arr.Delete(0)
	fmt.Println(arr.String())
	arr.Delete(0)
	fmt.Println(arr.String())
	arr.Delete(0)
	fmt.Println(arr.String())

	// 超范围插入
	defer func() {
		err := recover() // 捕获错误
		fmt.Println(err) // 打印错误信息
	}()
	arr.Insert(111, 2333)
	fmt.Println(arr.String())
}
