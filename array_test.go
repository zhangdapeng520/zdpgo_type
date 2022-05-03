package zdpgo_type

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_test"
	"testing"
)

var (
	T = zdpgo_test.New()
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

func TestArray_Index(t *testing.T) {
	arr := NewArray("a", "b", "c")
	tables := []struct {
		value string
		index int
	}{
		{"a", 0},
		{"b", 1},
		{"c", 2},
		{"d", -1},
	}
	T.SetTestObj(t)
	for _, table := range tables {
		tIndex := arr.Index(table.value)
		T.Assert.Equal(table.index, tIndex)
	}
}

func TestArray_Contains(t *testing.T) {
	arr := NewArray("a", "b", "c")
	tables := []struct {
		value string
		flag  bool
	}{
		{"a", true},
		{"b", true},
		{"c", true},
		{"d", false},
	}
	T.SetTestObj(t)
	for _, table := range tables {
		tFlag := arr.Contains(table.value)
		T.Assert.Equal(table.flag, tFlag)
	}
}

func TestArray_Delete(t *testing.T) {
	arr := NewArray("a", "b", "c")
	tables := []struct {
		index int
		flag  bool
	}{
		{2, false},
		{1, false},
		{0, false},
	}
	T.SetTestObj(t)
	for _, table := range tables {
		value := arr.Get(table.index)
		err := arr.Delete(table.index)
		T.Assert.Equal(err, nil)
		T.Assert.Equal(table.flag, arr.Contains(value))
	}
}
