package treeset

import (
	"encoding/json"

	"github.com/zhangdapeng520/zdpgo_type/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Set[string])(nil)
	var _ containers.JSONDeserializer = (*Set[string])(nil)
}

// ToJSON outputs the JSON representation of the set.
func (set *Set[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[T]) FromJSON(data []byte) error {
	elements := []T{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
