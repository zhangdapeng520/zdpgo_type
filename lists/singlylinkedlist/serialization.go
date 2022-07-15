package singlylinkedlist

import (
	"encoding/json"

	"github.com/zhangdapeng520/zdpgo_type/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*List[string])(nil)
	var _ containers.JSONDeserializer = (*List[string])(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	elements := []T{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
