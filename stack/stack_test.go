package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func getStack() *Stack {
	return NewStack()
}

// 测试栈的基本使用
func TestStack_basic(t *testing.T) {
	s := getStack()

	// 入栈
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.String())

	// 出栈
	fmt.Println(s.Pop())
	fmt.Println(s.values)
	fmt.Println(s.String())

	// 出栈并还原类型
	v := s.Pop()
	v1 := v.(int)
	fmt.Println(reflect.TypeOf(v), reflect.TypeOf(v1))
}

// 测试栈存储字符串
func TestStack_str(t *testing.T) {
	s := getStack()

	// 入栈
	s.Push("1")
	s.Push("2")
	s.Push("3")
	fmt.Println(s.String())

	// 出栈
	fmt.Println(s.Pop())
	fmt.Println(s.values)
	fmt.Println(s.String())

	// 出栈并还原类型
	v := s.Pop()
	v1 := v.(string)
	fmt.Println(reflect.TypeOf(v), reflect.TypeOf(v1))
}
