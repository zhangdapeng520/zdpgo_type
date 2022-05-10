package zdpgo_type

import "errors"

/*
@Time : 2022/5/10 15:07
@Author : 张大鹏
@File : array_bool.go
@Software: Goland2021.3.1
@Description: 布尔数组类型
*/

type ArrayBool struct {
	size   int
	values []bool
}

func NewArrayBool() *ArrayBool {
	return &ArrayBool{}
}

func (a *ArrayBool) Append(b bool) {
	a.values = append(a.values, b)
	a.size++
}

func (a *ArrayBool) Pop() (b bool, e error) {
	if a.size < 1 {
		e = errors.New("数组元素为空")
		return
	}
	b = a.values[a.size-1]
	a.values = a.values[:a.size-1]
	a.size--
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
