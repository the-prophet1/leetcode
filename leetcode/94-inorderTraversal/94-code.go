package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	InorderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	InorderTraversal(root.Right, result)
}
func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	InorderTraversal(root, &result)
	return result
}

func main() {
	rightLeft := &TreeNode{3, nil, nil}
	right := &TreeNode{2, rightLeft, nil}
	root := &TreeNode{1, nil, right}
	result := inorderTraversal(root)
	fmt.Println(result)
}
