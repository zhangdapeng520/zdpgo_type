package node

import (
	"fmt"
	"testing"
)

func getLinkedList() *LinkedList {
	return NewLinkedList()
}

// 链表的基本使用
func TestLinkedList_basic(t *testing.T) {
	l := getLinkedList()

	// 插入元素
	l.Insert(0, 1)
	l.Insert(1, 22)
	l.Insert(0, 11)
	l.Insert(3, 33)
	fmt.Println(l.String())

	// 根据索引获取元素的值
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))

	// 根据索引修改元素的值
	l.Set(1, 111)
	fmt.Println(l.Get(1))
	fmt.Println(l.String())

	// 追加元素
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l.String())

	// 在开头插入元素
	l.Insert(0, 1)
	l.Insert(0, 1)
	l.Insert(0, 1)
	fmt.Println(l.String())

	// 删除开头元素
	l.Delete(0)
	l.Delete(0)
	l.Delete(0)
	fmt.Println(l.String())

	// 删除最后的元素
	l.Delete(l.Size - 1)
	l.Delete(l.Size - 1)
	l.Delete(l.Size - 1)
	fmt.Println(l.String())

	// 删除中间的元素
	l.Delete(1)
	l.Delete(1)
	fmt.Println(l.String())
}
