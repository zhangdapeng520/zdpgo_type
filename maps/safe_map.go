package maps

import (
	"sync"
)

// SafeMap 线程安全的map类型
type SafeMap[K comparable, V any] struct {
	sync.RWMutex
	items StdMap[K, V]
}

// Clear 清空元素
func (m *SafeMap[K, V]) Clear() {
	if m.items == nil {
		return
	}
	m.Lock()
	m.items = nil
	m.Unlock()
}

// Set 设置键值对
func (m *SafeMap[K, V]) Set(k K, v V) {
	m.Lock()
	if m.items == nil {
		m.items = map[K]V{k: v}
	} else {
		m.items[k] = v
	}
	m.Unlock()
}

// Get 获取值
func (m *SafeMap[K, V]) Get(k K) (v V) {
	v, _ = m.Load(k)
	return
}

// Has 判断元素是否存在
func (m *SafeMap[K, V]) Has(k K) (exists bool) {
	_, exists = m.Load(k)
	return
}

// Load 加载值
func (m *SafeMap[K, V]) Load(k K) (v V, ok bool) {
	if m.items == nil {
		return
	}
	m.RLock()
	if m.items != nil {
		v, ok = m.items[k]
	}
	m.RUnlock()
	return
}

// Delete 删除键值对
func (m *SafeMap[K, V]) Delete(k K) {
	m.Lock()
	m.items.Delete(k)
	m.Unlock()
}

// Values 获取所有的值
func (m *SafeMap[K, V]) Values() (v []V) {
	if m.items == nil {
		return
	}
	m.RLock()
	v = m.items.Values()
	m.RUnlock()
	return
}

// Keys 获取所有的键
func (m *SafeMap[K, V]) Keys() (keys []K) {
	if m.items == nil {
		return nil
	}
	m.RLock()
	keys = m.items.Keys()
	m.RUnlock()
	return
}

// Len 获取元素个数
func (m *SafeMap[K, V]) Len() (l int) {
	if m.items == nil {
		return
	}
	m.RLock()
	l = m.items.Len()
	m.RUnlock()
	return
}

// Range 遍历map
func (m *SafeMap[K, V]) Range(f func(k K, v V) bool) {
	if m.items == nil {
		return
	}
	m.RLock()
	defer m.RUnlock()
	m.items.Range(f)
}

// Merge 合并两个map
func (m *SafeMap[K, V]) Merge(in MapI[K, V]) {
	if m.items == nil {
		m.items = make(map[K]V, in.Len())
	}
	m.Lock()
	defer m.Unlock()
	m.items.Merge(in)
}

// Equal 判断两个map是否相等
func (m *SafeMap[K, V]) Equal(m2 MapI[K, V]) bool {
	m.RLock()
	defer m.RUnlock()
	return m.items.Equal(m2)
}

// MarshalBinary 解析为字节数组
func (m *SafeMap[K, V]) MarshalBinary() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()
	return m.items.MarshalBinary()
}

// UnmarshalBinary 将自己数组加载为map
func (m *SafeMap[K, V]) UnmarshalBinary(data []byte) (err error) {
	m.Lock()
	defer m.Unlock()
	return m.items.UnmarshalBinary(data)
}

// MarshalJSON 实现json序列化接口
func (m *SafeMap[K, V]) MarshalJSON() (out []byte, err error) {
	m.RLock()
	defer m.RUnlock()
	return m.items.MarshalJSON()
}

// UnmarshalJSON 实现json序列化接口
func (m *SafeMap[K, V]) UnmarshalJSON(in []byte) (err error) {
	m.Lock()
	defer m.Unlock()
	return m.items.UnmarshalJSON(in)
}

// String 获取字符串表示
func (m *SafeMap[K, V]) String() string {
	m.RLock()
	defer m.RUnlock()
	return m.items.String()
}
