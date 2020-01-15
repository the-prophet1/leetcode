package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func search(root *TreeNode, level int, data *[][]int) {
	if root == nil {
		return
	}
	if len(*data) < level+1 {
		*data = append(*data, []int{})
	}
	(*data)[level] = append((*data)[level], root.Val)
	search(root.Left, level+1, data)
	search(root.Right, level+1, data)
}

func levelOrder(root *TreeNode) [][]int {
	levelNodes := make([][]int, 0)
	search(root, 0, &levelNodes)
	return levelNodes
}

func main() {
	node5 := &TreeNode{2, nil, nil}
	node4 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{1, nil, node5}
	node2 := &TreeNode{1, node4, nil}
	node1 := &TreeNode{0, node2, node3}
	levelOrder(node1)
}
