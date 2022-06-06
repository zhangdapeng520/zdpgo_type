package main

import lls "github.com/zhangdapeng520/zdpgo_type/stacks/linkedliststack"

// LinkedListStackExample to demonstrate basic usage of LinkedListStack
func main() {
	stack := lls.New[int]() // empty
	stack.Push(1)           // 1
	stack.Push(2)           // 1, 2
	stack.Values()          // 2, 1 (LIFO order)
	_, _ = stack.Peek()     // 2,true
	_, _ = stack.Pop()      // 2, true
	_, _ = stack.Pop()      // 1, true
	_, _ = stack.Pop()      // nil, false (nothing to pop)
	stack.Push(1)           // 1
	stack.Clear()           // empty
	stack.Empty()           // true
	stack.Size()            // 0
}
