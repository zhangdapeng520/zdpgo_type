package node

import (
	"fmt"
	"sync"
)

// LinkedList 链表
type LinkedList struct {
	Size int        // 元素个数
	Head *Node      // 头节点
	Last *Node      // 尾节点
	lock sync.Mutex // 同步锁
}

func NewLinkedList() *LinkedList {
	l := LinkedList{}
	return &l
}

// Get 根据索引获取元素的值
func (l *LinkedList) Get(index int) interface{} {
	p := l.getNode(index)
	return p.Data
}

// Set 根据索引修改元素的值
func (l *LinkedList) Set(index int, value interface{}) {
	p := l.getNode(index)
	p.Data = value
}

// 获取指定索引的节点
func (l *LinkedList) getNode(index int) *Node {
	// 判断索引越界
	if index < 0 || index >= l.Size {
		panic("索引越界：给出的索引超过链表范围！")
	}

	// 创建一个节点，指向头节点
	p := l.Head

	// 一直往下找
	for i := 0; i < index; i++ {
		p = p.Next
	}

	// 找到了返回该节点
	return p
}

// Insert 插入数据
func (l *LinkedList) Insert(index int, data interface{}) {
	// 判断索引越界
	if index < 0 || index > l.Size {
		panic("索引越界：给出的索引超过链表范围！")
	}

	l.lock.Lock()
	// 创建一个节点
	node := &Node{Data: data}

	if l.Size == 0 { // 空链表
		l.Head = node
		l.Last = node
		l.Head.Next = node
	} else if index == 0 { // 不是空链表，且插入头部
		// 注意：这里有点绕，想象成同时存在两个链表，在交换
		node.Next = l.Head // 当前节点的下一个节点为原本链表的头节点
		l.Head = node      // 链表的头节点为当前节点
	} else if index == l.Size { // 插入尾部
		l.Last.Next = node // 链表的最后一个节点的下一个节点为新增的节点
		l.Last = node      // 链表的最后一个节点为新增的节点
	} else { // 插入中间
		prevNode := l.getNode(index - 1) // 获取前置节点
		node.Next = prevNode.Next        // 当前节点的下一个节点为指定索引前一个节点的下一个节点
		prevNode.Next = node             // 前一个节点的下一个节点为当前节点
	}
	l.Size += 1
	l.lock.Unlock()
}

// Append 在末尾追加元素
func (l *LinkedList) Append(data interface{}) {
	l.Insert(l.Size, data)
}

// Delete 删除指定索引的元素
func (l *LinkedList) Delete(index int) *Node {
	// 判断索引越界
	if index < 0 || index >= l.Size {
		panic("索引越界：给出的索引超过链表范围！")
	}

	l.lock.Lock()
	removedNode := &Node{}
	if index == 0 { // 删除头节点
		removedNode = l.Head // 删除的是头节点
		l.Head = l.Head.Next // 将头节点设置为头节点的下一个节点，实现删除
	} else if index == l.Size-1 { // 插入尾部
		prevNode := l.getNode(index - 1) // 获取前置节点
		removedNode = prevNode.Next      // 要删除的节点是上一个节点的下一个节点，也就是当前节点
		prevNode.Next = nil              // 上一个节点的下一个节点为空，因为被删除了
		l.Last = prevNode                // 最后一个节点为上一个节点
	} else { // 插入中间
		prevNode := l.getNode(index - 1) // 获取前置节点
		nextNode := prevNode.Next.Next   // 获取后置节点，是上一个节点的下一个节点的下一个节点
		removedNode = prevNode.Next      // 要删除的节点是上一个节点的下一个节点，也就是当前节点
		prevNode.Next = nextNode         //删除了中间的节点
	}
	l.Size -= 1
	l.lock.Unlock()

	// 返回被删除的节点
	return removedNode
}

// String 获取数组的字符串表示
func (l *LinkedList) String() string {
	l.lock.Lock()
	out := "["
	p := l.Head
	for {
		out += fmt.Sprintf("%v", p.Data)
		if p.Next == nil {
			break
		} else {
			out += ","
		}
		p = p.Next
	}
	out += "]"
	l.lock.Unlock()
	return out
}
