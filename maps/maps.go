package maps

import "github.com/zhangdapeng520/zdpgo_type/containers"

// Map 所有map类型都应该实现的方法
type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[K comparable, V any] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
