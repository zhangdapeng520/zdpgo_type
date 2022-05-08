package zdpgo_type

import (
	"fmt"
	"testing"
)

// 测试数组的基本使用
func TestArray_basic(t *testing.T) {
	arr := NewArray(1, 2, 3, 4)
	fmt.Println(arr.ToString())

	// 根据索引读取元素
	fmt.Println(arr.Get(0), arr.Get(1))

	// 更新元素
	arr.Set(0, 33)
	fmt.Println(arr.ToString())

	// 尾部插入
	arr.Append(111)
	fmt.Println(arr.ToString())

	// 中间插入
	arr.Insert(1, 44)
	fmt.Println(arr.ToString())

	// 头部插入
	arr.Insert(0, 100)
	fmt.Println(arr.ToString())

	// 删除元素
	arr.Delete(0)
	fmt.Println(arr.ToString())
	arr.Delete(0)
	fmt.Println(arr.ToString())
	arr.Delete(0)
	fmt.Println(arr.ToString())

	// 超范围插入
	defer func() {
		err := recover() // 捕获错误
		fmt.Println(err) // 打印错误信息
	}()
	arr.Insert(111, 2333)
	fmt.Println(arr.ToString())
}

func TestArray_Remove(t *testing.T) {
	arr := NewArray("a", "b", "c")
	err := arr.Remove("a")
	fmt.Println(arr.ToString(), err)
	fmt.Println(arr.ToSliceValues(), err)

	err = arr.Remove("a")
	fmt.Println(arr.ToString(), err)

	err = arr.Remove("b")
	fmt.Println(arr.ToString(), err)

	err = arr.Remove("c")
	fmt.Println(arr.ToString(), err)

	err = arr.Remove("d")
	fmt.Println(arr.ToString(), err)
}

func TestArray_IsExists(t *testing.T) {
	// 字符串
	arr := NewArray("a", "b", "c")
	fmt.Println(arr.IsExists("a"))
	fmt.Println(arr.IsExists("aa"))

	// 数字
	arr1 := NewArray(1, 2, 3)
	fmt.Println(arr1.IsExists(1))
	fmt.Println(arr1.IsExists(11))
}

func TestArray_AddIfNotExists(t *testing.T) {
	arr := NewArray("a", "b", "c")
	arr.AddIfNotExists("a")
	arr.AddIfNotExists("aa")
	arr.AddIfNotExists("b")
	arr.AddIfNotExists("bb")
	arr.AddIfNotExists("c")
	arr.AddIfNotExists("cc")
	fmt.Println(arr.ToSliceValues())
}
