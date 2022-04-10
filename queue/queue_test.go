package queue

import (
	"fmt"
	"testing"
)

func getQueue() *Queue {
	return NewQueue(10)
}

// 测试队列的基本使用
func TestQueue_basic(t *testing.T) {
	q := getQueue()

	// 入队
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.String())

	// 出队
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.String())
	fmt.Println(q.Dequeue())
	fmt.Println(q.String())
}
