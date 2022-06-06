package main

import "github.com/zhangdapeng520/zdpgo_type/stacks/arraystack"

// ArrayStackExample to demonstrate basic usage of ArrayStack
func main() {
	stack := arraystack.New[int]() // empty
	stack.Push(1)                  // 1
	stack.Push(2)                  // 1, 2
	stack.Values()                 // 2, 1 (LIFO order)
	_, _ = stack.Peek()            // 2,true
	_, _ = stack.Pop()             // 2, true
	_, _ = stack.Pop()             // 1, true
	_, _ = stack.Pop()             // nil, false (nothing to pop)
	stack.Push(1)                  // 1
	stack.Clear()                  // empty
	stack.Empty()                  // true
	stack.Size()                   // 0
}
