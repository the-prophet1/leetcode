package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先遍历方式
func maxDepth(root *TreeNode) int {
	if root == nil { //root为空时，深度为0
		return 0
	}
	leftDepth := maxDepth(root.Left)   //递归获取左子树的最大深度
	RightDepth := maxDepth(root.Right) //递归获取右子树的最大深度

	return int(math.Max(float64(leftDepth), float64(RightDepth)) + 1) //取左子树和右子树最大深度的较大值并加1
}

func search(root *TreeNode, level int, maxLevel *int) {
	if root == nil {
		return
	}
	if *maxLevel < level+1 { //当前深度大于最大深度
		*maxLevel = level + 1
	}
	search(root.Left, level+1, maxLevel)
	search(root.Right, level+1, maxLevel)
}

func maxDepth1(root *TreeNode) int {
	maxLevel := 0
	search(root, 0, &maxLevel)
	return maxLevel
}

func main() {
	//	node5 := &TreeNode{2,nil,nil}
	node4 := &TreeNode{2, nil, nil}
	//	node3 := &TreeNode{1,nil,node5}
	node2 := &TreeNode{1, node4, nil}
	node1 := &TreeNode{0, node2, nil}
	fmt.Println(maxDepth(node1))
	fmt.Println(maxDepth1(node1))
}
