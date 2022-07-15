package trees

import "github.com/zhangdapeng520/zdpgo_type/containers"

// Tree interface that all trees implement
type Tree[T comparable] interface {
	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
