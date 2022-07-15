package zdpgo_type

import (
	"sync"
)

// Set 实现 set 集合，变相实现 切片去重
type Set[T comparable] struct {
	data         map[T]bool
	sync.RWMutex // 不用变量接收，则通过对象调用
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: map[T]bool{},
	}
}

// Add 添加元素到集合
func (s *Set[T]) Add(items ...T) {
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
func (s *Set[T]) Remove(items ...T) {
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
func (s *Set[T]) Contains(item T) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.data[item]
	return ok
}

// Length 获取集合的元素个数
func (s *Set[T]) Length() int {
	return len(s.data)
}

// Clear 清空集合元素
func (s *Set[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data = map[T]bool{}
}

// IsEmpty 判断集合是否为空
func (s *Set[T]) IsEmpty() bool {
	return s.Length() == 0
}

// ToSlice 转换为切片
func (s *Set[T]) ToSlice() []T {
	s.RLock()
	defer s.RUnlock()
	var list []T
	for item := range s.data {
		list = append(list, item)
	}
	return list
}
