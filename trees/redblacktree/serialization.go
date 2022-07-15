package redblacktree

import (
	"encoding/json"

	"github.com/zhangdapeng520/zdpgo_type/containers"
	"github.com/zhangdapeng520/zdpgo_type/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Tree[string, string])(nil)
	var _ containers.JSONDeserializer = (*Tree[string, string])(nil)
}

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[string, V]) FromJSON(data []byte) error {
	elements := make(map[string]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
