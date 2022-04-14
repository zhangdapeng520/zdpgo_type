package tree

import (
	"fmt"
	"testing"
)

func TestTreeNode_New(t *testing.T) {
	// 创建一棵树
	root := CreateTreeNode(5)
	root.Left = CreateTreeNode(2)
	root.Right = CreateTreeNode(4)
	root.Left.Right = CreateTreeNode(7)
	root.Left.Right.Left = CreateTreeNode(6)
	root.Right.Left = CreateTreeNode(8)
	root.Right.Right = CreateTreeNode(9)

	fmt.Printf("%d\n", root.FindNode(root, 4).Data)
	fmt.Printf("%d\n", root.GetTreeHeight(root))

	root.PreOrder(root)
	fmt.Printf("\n")
	root.InOrder(root)
	fmt.Printf("\n")
	root.PostOrder(root)
	fmt.Printf("\n")

	root.GetLeafNode(root)
	fmt.Printf("\n")
}
