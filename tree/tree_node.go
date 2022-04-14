package tree

import "fmt"

// TreeNode 二叉树的节点
type TreeNode struct {
	Data  interface{} // 数据
	Left  *TreeNode   // 左节点
	Right *TreeNode   // 右节点
}

// Print 打印节点的值
func (node *TreeNode) Print() {
	fmt.Printf("%d ", node.Data)
}

// SetValue 设置节点的值
func (node *TreeNode) SetValue(value int) {
	node.Data = value
}

// CreateTreeNode 创建节点
func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

// FindNode 查找节点，利用递归进行查找
func (node *TreeNode) FindNode(n *TreeNode, x int) *TreeNode {
	if n == nil {
		return nil
	} else if n.Data == x {
		return n
	} else {
		p := node.FindNode(n.Left, x)
		if p != nil {
			return p
		}
		return node.FindNode(n.Right, x)
	}
}

// GetTreeHeight 求树的高度：树的高度，树的高度=Max(左子树高度，右子树高度)+1
func (node *TreeNode) GetTreeHeight(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		lHeight := node.GetTreeHeight(n.Left)
		rHeight := node.GetTreeHeight(n.Right)
		if lHeight > rHeight {
			return lHeight + 1
		} else {
			return rHeight + 1
		}
	}
}

// PreOrder 递归前序遍历二叉树
func (node *TreeNode) PreOrder(n *TreeNode) {
	if n != nil {
		fmt.Printf("%d ", n.Data)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

// InOrder 递归中序遍历二叉树
func (node *TreeNode) InOrder(n *TreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		fmt.Printf("%d ", n.Data)
		node.InOrder(n.Right)
	}
}

// PostOrder 递归后序遍历二叉树
func (node *TreeNode) PostOrder(n *TreeNode) {
	if n != nil {
		node.PostOrder(n.Left)
		node.PostOrder(n.Right)
		fmt.Printf("%d ", n.Data)
	}
}

// GetLeafNode 打印所有的叶子节点
func (node *TreeNode) GetLeafNode(n *TreeNode) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			fmt.Printf("%d ", n.Data)
		}
		node.GetLeafNode(n.Left)
		node.GetLeafNode(n.Right)
	}
}
