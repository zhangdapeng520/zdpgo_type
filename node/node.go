package node

type Node struct {
	Data int   // 数据
	Next *Node // 下一个节点
}

// NewNode 新建一个节点
func NewNode() *Node {
	n := Node{}
	return &n
}
