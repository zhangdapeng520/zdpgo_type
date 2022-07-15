package treebidimap

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_type/maps"
	"github.com/zhangdapeng520/zdpgo_type/trees/redblacktree"
	"github.com/zhangdapeng520/zdpgo_type/utils"
)

func assertMapImplementation() {
	var _ maps.BidiMap[string, string] = (*Map[string, string])(nil)
}

// Map holds the elements in two red-black trees.
type Map[K comparable, V comparable] struct {
	forwardMap      redblacktree.Tree[K, *data[K, V]]
	inverseMap      redblacktree.Tree[V, *data[K, V]]
	keyComparator   utils.Comparator[K]
	valueComparator utils.Comparator[V]
}

type data[K comparable, V comparable] struct {
	key   K
	value V
}

// NewWith instantiates a bidirectional maps.
func NewWith[K comparable, V comparable](keyComparator utils.Comparator[K],
	valueComparator utils.Comparator[V]) *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      *redblacktree.NewWith[K, *data[K, V]](keyComparator),
		inverseMap:      *redblacktree.NewWith[V, *data[K, V]](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithIntComparators instantiates a bidirectional maps with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithIntComparators() *Map[int, int] {
	return NewWith(utils.NumberComparator[int], utils.NumberComparator[int])
}

// NewWithStringComparators instantiates a bidirectional maps with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators() *Map[string, string] {
	return NewWith(utils.StringComparator, utils.StringComparator)
}

// Put inserts element into the maps.
func (m *Map[K, V]) Put(key K, value V) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d.value)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d.key)
	}
	d := &data[K, V]{key: key, value: value}
	m.forwardMap.Put(key, d)
	m.inverseMap.Put(value, d)
}

// Get searches the element in the maps by key and returns its value or nil if key is not found in maps.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d.value, true
	}
	return *new(V), false
}

// GetKey searches the element in the maps by value and returns its key or nil if value is not found in maps.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d.key, true
	}
	return *new(K), false
}

// Remove removes the element from the maps by key.
func (m *Map[K, V]) Remove(key K) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d.value)
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

// Keys returns all keys (ordered).
func (m *Map[K, V]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
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
	str := "TreeBidiMap\nmaps["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
