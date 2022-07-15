package hashbidimap

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_type/maps"
	"github.com/zhangdapeng520/zdpgo_type/maps/hashmap"
)

func assertMapImplementation() {
	var _ maps.BidiMap[string, string] = (*Map[string, string])(nil)
}

// Map holds the elements in two hashmaps.
type Map[K comparable, V comparable] struct {
	forwardMap hashmap.Map[K, V]
	inverseMap hashmap.Map[V, K]
}

// New instantiates a bidirectional maps.
func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{*hashmap.New[K, V](), *hashmap.New[V, K]()}
}

// Put inserts element into the maps.
func (m *Map[K, V]) Put(key K, value V) {
	if valueByKey, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(valueByKey)
	}
	if keyByValue, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(keyByValue)
	}
	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the maps by key and returns its value or nil if key is not found in maps.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.forwardMap.Get(key)
}

// GetKey searches the element in the maps by value and returns its key or nil if value is not found in maps.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	return m.inverseMap.Get(value)
}

// Remove removes the element from the maps by key.
func (m *Map[K, V]) Remove(key K) {
	if value, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(value)
	}
}

// Empty returns true if maps does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the maps.
func (m *Map[K, V]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (random order).
func (m *Map[K, V]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (random order).
func (m *Map[K, V]) Values() []V {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the maps.
func (m *Map[K, V]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "HashBidiMap\n"
	str += fmt.Sprintf("%v", m.forwardMap)
	return str
}
