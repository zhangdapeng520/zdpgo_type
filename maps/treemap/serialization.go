package treemap

import "github.com/zhangdapeng520/zdpgo_type/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map[string, string])(nil)
	var _ containers.JSONDeserializer = (*Map[string, string])(nil)
}

// ToJSON outputs the JSON representation of the maps.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the maps from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}
