package queue

import (
	"fmt"
	"sync"
)

// Queue 队列，先进先出
type Queue struct {
	values   []interface{} // 数据内容
	Capacity int           // 容量
	Size     int           // 元素个数
	lock     sync.Mutex    // 同步锁
}

func NewQueue(capacity int) *Queue {
	q := Queue{
		Capacity: capacity,
	}
	return &q
}

// Enqueue 入队操作
func (q *Queue) Enqueue(value interface{}) {
	// 判断是否已满
	if q.Size >= q.Capacity {
		panic("队列已满！")
	}

	// 入队
	q.lock.Lock()
	q.values = append(q.values, value)
	q.Size += 1
	q.lock.Unlock()
}

// Dequeue 出队操作
func (q *Queue) Dequeue() interface{} {
	if q.Size == 0 {
		panic("队列已空！")
	}

	// 出队的是队首元素，先进先出
	q.lock.Lock()
	dequeueElement := q.values[0]
	q.values = q.values[1:]
	q.Size -= 1
	q.lock.Unlock()

	// 返回出队元素
	return dequeueElement
}

// String 返回队列的字符串表示
func (q *Queue) String() string {
	q.lock.Lock()
	out := "["
	arrLength := q.Size - 1
	for i, v := range q.values {
		out += fmt.Sprintf("%v", v)
		if i != arrLength {
			out += ","
		}
	}
	out += "]"
	q.lock.Unlock()
	return out
}
