package stack

import (
	"fmt"
	"sync"
)

// Stack 栈，先进后出FILO
type Stack struct {
	values []interface{} // 元素内容
	Size   int           // 元素个数
	lock   sync.Mutex    // 同步锁
}

func NewStack() *Stack {
	s := Stack{}
	return &s
}

// Push 入栈
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	s.values = append(s.values, v)
	s.Size += 1
	s.lock.Unlock()
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	lastIndex := len(s.values) - 1
	lastValue := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	s.Size -= 1
	s.lock.Unlock()
	return lastValue
}

// String 获取栈的字符串表示
func (s *Stack) String() string {
	s.lock.Lock()
	out := "["
	for i, v := range s.values {
		vs := fmt.Sprintf("%v", v)
		out += vs
		if i < s.Size-1 {
			out += ","
		}
	}
	out += "]"
	s.lock.Unlock()
	return out
}
