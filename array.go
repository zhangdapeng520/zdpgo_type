package zdpgo_type

import (
	"errors"
	"fmt"
	"sync"
)

// Array 给整数切片取个别名
type Array[V comparable] struct {
	lock   sync.Mutex // 保证线程安全
	Size   int        // 元素个数
	values []V
}

// NewArray 创建数组
func NewArray[V comparable](values ...V) *Array[V] {
	arr := Array[V]{}
	arr.values = values
	arr.Size = len(values)
	return &arr
}

// Get 根据索引获取元素
func (arr *Array[V]) Get(index int) V {
	arr.lock.Lock()
	value := arr.values[index]
	arr.lock.Unlock()
	return value
}

// Set 根据索引修改元素
func (arr *Array[V]) Set(index int, value V) {
	arr.lock.Lock()
	arr.values[index] = value
	arr.lock.Unlock()
}

// Append 尾部插入
func (arr *Array[V]) Append(value V) {
	arr.lock.Lock()
	arr.values = append(arr.values, value)
	arr.Size += 1
	arr.lock.Unlock()
}

// Insert 中间插入
func (arr *Array[V]) Insert(index int, value V) {
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
func (arr *Array[V]) Delete(index int) {
	if index < 0 || index > arr.Size {
		panic("索引越界：数组不存在该索引！")
	}
	arr.lock.Lock()
	tail := arr.values[index+1:]                     // 剩下的元素
	arr.values = append(arr.values[:index], tail...) // 插入剩下的元素
	arr.Size -= 1
	arr.lock.Unlock()
}

// Pop 删除并返回最后一个元素
func (arr *Array[V]) Pop() (v V) {
	// 特殊情况
	if arr.Size == 0 {
		return
	}

	// 获取最后一个元素
	index := arr.Size - 1
	value := arr.values[index]

	// 删除最后一个元素
	arr.values = arr.values[:index]

	// 返回
	return value
}

// PopFirst 删除并返回第一个元素
func (arr *Array[V]) PopFirst() (v V) {
	// 特殊情况
	if arr.Size == 0 {
		return
	}

	// 获取最后一个元素
	value := arr.values[0]

	// 删除最后一个元素
	arr.values = arr.values[1:]

	// 返回
	return value
}

func (arr *Array[V]) Remove(element V) (err error) {
	arr.lock.Lock()
	if arr.Size == 0 {
		err = errors.New("数组为空，无法删除元素")
		arr.lock.Unlock()
		return
	}

	// 遍历数组，找到该元素然后删除
	var newArr []V
	var flag bool
	for _, vv := range arr.values {
		if vv != element {
			newArr = append(newArr, vv)
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

// ToString 获取数组的字符串表示
func (arr *Array[V]) ToString() string {
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

// ToSliceValues 获取数组的切片表示
func (arr *Array[V]) ToSliceValues() []V {
	return arr.values
}

// Length 获取数组的长度
func (arr *Array[V]) Length() int {
	return arr.Size
}

// IsExists 判断元素是否已存在
func (arr *Array[V]) IsExists(element V) bool {
	for _, v := range arr.values {
		if v == element {
			return true
		}
	}
	return false
}

// AddIfNotExists 如果不存在 ，则添加元素
func (arr *Array[V]) AddIfNotExists(element V) {
	if !arr.IsExists(element) {
		arr.Append(element)
	}
}
