package linkedhashmap

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_type/maps"
	"strings"

	"github.com/zhangdapeng520/zdpgo_type/lists/doublylinkedlist"
)

func assertMapImplementation() {
	var _ maps.Map[string, string] = (*Map[string, string])(nil)
}

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[K comparable, V comparable] struct {
	table    map[K]V
	ordering *doublylinkedlist.List[K]
}

// New instantiates a linked-hash-maps.
func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{
		table:    make(map[K]V),
		ordering: doublylinkedlist.New[K](),
	}
}

// Put inserts key-value pair into the maps.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

// Get searches the element in the maps by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.table[key]
	return
}

// Remove removes the element from the maps by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Empty returns true if maps does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the maps.
func (m *Map[K, V]) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.ordering.Values()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the maps.
func (m *Map[K, V]) Clear() {
	m.table = make(map[K]V)
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "LinkedHashMap\nmaps["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
