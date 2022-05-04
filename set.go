package zdpgo_type

import (
	"sync"
)

/*
@Time : 2022/5/4 21:32
@Author : 张大鹏
@File : set
@Software: Goland2021.3.1
@Description: Set集合类型
*/

// Set 实现 set 集合，变相实现 切片去重
type Set struct {
	data         map[interface{}]bool
	sync.RWMutex // 不用变量接收，则通过对象调用
}

func NewSet() *Set {
	return &Set{
		data: map[interface{}]bool{},
	}
}

// Add 添加元素到集合
func (s *Set) Add(items ...interface{}) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.data[item] = true
	}
}

// Remove 从集合中移除元素
func (s *Set) Remove(items ...interface{}) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(s.data, item)
	}
}

// Contains 判断集合是否拥有某个元素
func (s *Set) Contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.data[item]
	return ok
}

// Length 获取集合的元素个数
func (s *Set) Length() int {
	return len(s.data)
}

// Clear 清空集合元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data = map[interface{}]bool{}
}

// IsEmpty 判断集合是否为空
func (s *Set) IsEmpty() bool {
	return s.Length() == 0
}

// ToSlice 转换为切片
func (s *Set) ToSlice() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for item := range s.data {
		list = append(list, item)
	}
	return list
}
