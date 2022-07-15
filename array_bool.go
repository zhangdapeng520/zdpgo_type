package zdpgo_type

import (
	"errors"
	"sync"
)

type ArrayBool struct {
	size   int
	values []bool
	sync.Mutex
}

func NewArrayBool(values ...bool) *ArrayBool {
	a := &ArrayBool{}
	for _, v := range values {
		a.values = append(a.values, v)
	}
	return a
}

func (a *ArrayBool) Append(b bool) {
	a.Lock()
	a.values = append(a.values, b)
	a.size++
	a.Unlock()
}

func (a *ArrayBool) Pop() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	a.Lock()
	b = a.values[a.size-1]
	a.values = a.values[:a.size-1]
	a.size--
	a.Unlock()
	return
}

func (a *ArrayBool) IsAllTrue() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	for _, v := range a.values {
		if !v {
			return
		}
	}
	b = true
	return
}

// IsAnyTrue 是否有一个为true
func (a *ArrayBool) IsAnyTrue() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	for _, v := range a.values {
		if v {
			b = true
			return
		}
	}
	return
}

// IsAnyFalse 是否有一个为false
func (a *ArrayBool) IsAnyFalse() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	for _, v := range a.values {
		if !v {
			b = true
			return
		}
	}
	return
}

func (a *ArrayBool) IsAllFalse() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	for _, v := range a.values {
		if v {
			b = true
			return
		}
	}
	return
}

// Length 获取元素个数
func (a *ArrayBool) Length() int {
	return a.size
}

// ToSlice 转换为切片
func (a *ArrayBool) ToSlice() []bool {
	return a.values
}
