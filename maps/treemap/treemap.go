package treemap

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_type/maps"

	rbt "github.com/zhangdapeng520/zdpgo_type/trees/redblacktree"
	"github.com/zhangdapeng520/zdpgo_type/utils"
)

func assertMapImplementation() {
	var _ maps.Map[string, string] = (*Map[string, string])(nil)
}

// Map holds the elements in a red-black tree
type Map[K comparable, V comparable] struct {
	tree *rbt.Tree[K, V]
}

// NewWith instantiates a tree maps with the custom comparator.
func NewWith[K comparable, V comparable](comparator utils.Comparator[K]) *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWith[K, V](comparator)}
}

// NewWithIntComparator instantiates a tree maps with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[V comparable]() *Map[int, V] {
	return &Map[int, V]{tree: rbt.NewWithIntComparator[V]()}
}

// NewWithStringComparator instantiates a tree maps with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[V comparable]() *Map[string, V] {
	return &Map[string, V]{tree: rbt.NewWithStringComparator[V]()}
}

// Put inserts key-value pair into the maps.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

// Get searches the element in the maps by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the maps by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

// Empty returns true if maps does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the maps.
func (m *Map[K, V]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	return m.tree.Values()
}

// Clear removes all elements from the maps.
func (m *Map[K, V]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree maps.
// Returns nil, nil if maps is empty.
func (m *Map[K, V]) Min() (key K, value V) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	return *new(K), *new(V)
}

// Max returns the maximum key and its value from the tree maps.
// Returns nil, nil if maps is empty.
func (m *Map[K, V]) Max() (key K, value V) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	return *new(K), *new(V)
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the maps is empty, or because
// all keys in the maps are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value
	}
	return *new(K), *new(V)
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the maps is empty, or because
// all keys in the maps are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value
	}
	return *new(K), *new(V)
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "TreeMap\nmaps["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
