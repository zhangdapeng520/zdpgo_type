package zdpgo_type

import "testing"

func TestQueue_basic(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.IsEmpty())
}
