package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right != nil {
		return false
	}

	flag1 := isEqual(left.Left, right.Right)
	flag2 := isEqual(left.Right, right.Left)

	return flag1 && flag2 && left.Val == right.Val

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func main() {
	node5 := &TreeNode{2, nil, nil}
	node4 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{1, nil, node5}
	node2 := &TreeNode{1, node4, nil}
	node1 := &TreeNode{0, node2, node3}

	fmt.Println(isSymmetric(node1))
}
