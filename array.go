package zdpgo_type

import (
	"errors"
	"fmt"
	"sync"
)

// Array 给整数切片取个别名
type Array struct {
	lock   sync.Mutex // 保证线程安全
	Size   int        // 元素个数
	values []interface{}
}

// NewArray 创建数组
func NewArray(values ...interface{}) *Array {
	arr := Array{}
	arr.values = values
	arr.Size = len(values)
	return &arr
}

// Get 根据索引获取元素
func (arr *Array) Get(index int) interface{} {
	arr.lock.Lock()
	value := arr.values[index]
	arr.lock.Unlock()
	return value
}

// Set 根据索引修改元素
func (arr *Array) Set(index int, value interface{}) {
	arr.lock.Lock()
	arr.values[index] = value
	arr.lock.Unlock()
}

// Append 尾部插入
func (arr *Array) Append(value interface{}) {
	arr.lock.Lock()
	arr.values = append(arr.values, value)
	arr.Size += 1
	arr.lock.Unlock()
}

// Insert 中间插入
func (arr *Array) Insert(index int, value interface{}) {
	if index < 0 || index > arr.Size {
		panic("索引越界：数组不存在该索引！")
	}
	arr.lock.Lock()
	temp := arr.values[index]
	tail := arr.values[index:]                     // 剩下的元素
	arr.values = append(arr.values[:index], value) // 中间插入的元素
	arr.values = append(arr.values, tail...)       // 插入剩下的元素
	arr.values[index+1] = temp                     // 需要将tail被修改的值改回来
	arr.Size += 1
	arr.lock.Unlock()
}

// Delete 根据索引删除元素
func (arr *Array) Delete(index int) {
	if index < 0 || index > arr.Size {
		panic("索引越界：数组不存在该索引！")
	}
	arr.lock.Lock()
	tail := arr.values[index+1:]                     // 剩下的元素
	arr.values = append(arr.values[:index], tail...) // 插入剩下的元素
	arr.Size -= 1
	arr.lock.Unlock()
}

func (arr *Array) Remove(element interface{}) (err error) {
	arr.lock.Lock()
	if arr.Size == 0 {
		err = errors.New("数组为空，无法删除元素")
		arr.lock.Unlock()
		return
	}

	// 遍历数组，找到该元素然后删除
	var newArr []interface{}
	var flag bool
	for _, v := range arr.values {
		if v != element {
			newArr = append(newArr, v)
		} else {
			flag = true
		}
	}
	if !flag {
		err = errors.New("要删除的元素不存在")
		arr.lock.Unlock()
		return
	}

	// 替换表原本的数组
	arr.values = newArr
	arr.Size -= 1
	arr.lock.Unlock()

	return
}

// String 获取数组的字符串表示
func (arr *Array) String() string {
	arr.lock.Lock()
	out := "["
	arrLength := arr.Size - 1
	for i, v := range arr.values {
		out += fmt.Sprintf("%v", v)
		if i != arrLength {
			out += ","
		}
	}
	out += "]"
	arr.lock.Unlock()
	return out
}

// Length 获取数组的长度
func (arr *Array) Length() int {
	return arr.Size
}
