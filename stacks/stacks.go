package stacks

import "github.com/zhangdapeng520/zdpgo_type/containers"

// Stack interface that all stacks implement
type Stack[T comparable] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
