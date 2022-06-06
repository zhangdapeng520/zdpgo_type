package maps

// Map 定义map字典类型，key必须是可序列化的，v可以是任意类型
type Map[K comparable, V any] struct {
	items StdMap[K, V] // 标准map
}

// Clear 清空map的元素
func (m *Map[K, V]) Clear() {
	m.items = nil
}

// Len 获取元素的个数
func (m Map[K, V]) Len() int {
	return m.items.Len()
}

// Range 遍历集合元素
func (m Map[K, V]) Range(f func(k K, v V) bool) {
	m.items.Range(f)
}

// Load 根据键获取值
func (m Map[K, V]) Load(k K) (V, bool) {
	return m.items.Load(k)
}

// Get 根据键获取值
func (m Map[K, V]) Get(k K) V {
	return m.items.Get(k)
}

// Has 判断key是否存在
func (m Map[K, V]) Has(k K) bool {
	return m.items.Has(k)
}

// Delete 根据键删除键值对
func (m Map[K, V]) Delete(k K) {
	m.items.Delete(k)
}

// Keys 返回所有的key
func (m Map[K, V]) Keys() []K {
	return m.items.Keys()
}

// Values 返回所有的值
func (m Map[K, V]) Values() []V {
	return m.items.Values()
}

// Set 设置键值对
func (m *Map[K, V]) Set(k K, v V) {
	if m.items == nil {
		m.items = map[K]V{k: v}
	} else {
		m.items.Set(k, v)
	}
}

// Merge 合并两个map
func (m *Map[K, V]) Merge(in MapI[K, V]) {
	if m.items == nil {
		m.items = make(map[K]V, in.Len())
	}
	m.items.Merge(in)
}

// Equal 判断两个map的键值对是否完全一样
func (m Map[K, V]) Equal(m2 MapI[K, V]) bool {
	return m.items.Equal(m2)
}

// MarshalBinary 转换为字节数组
func (m Map[K, V]) MarshalBinary() ([]byte, error) {
	return m.items.MarshalBinary()
}

// UnmarshalBinary 将字节数组转换为map
func (m *Map[K, V]) UnmarshalBinary(data []byte) (err error) {
	return m.items.UnmarshalBinary(data)
}

// MarshalJSON 转换为json数据，实现此方法是为了支持json序列化
func (m Map[K, V]) MarshalJSON() (out []byte, err error) {
	return m.items.MarshalJSON()
}

// UnmarshalJSON 解析json数据，实现此方法是为了支持json序列化
func (m *Map[K, V]) UnmarshalJSON(in []byte) (err error) {
	return m.items.UnmarshalJSON(in)
}

// String 获取字符串表示
func (m Map[K, V]) String() string {
	return m.items.String()
}
