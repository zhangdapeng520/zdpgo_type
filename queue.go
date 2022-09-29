package zdpgo_type

// Queue 定义队列类型
// 队列特点：先进先出
type Queue[T any] []T

// NewQueue 创建队列
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push 追加元素
func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

// Pop 弹出元素
func (q *Queue[T]) Pop() T {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty 判断元素是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
