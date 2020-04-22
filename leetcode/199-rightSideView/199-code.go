package main

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

思路：
	对二叉树使用广度优先遍历,记录每一层的最后一个遍历的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0, 16)
	result := make([]int, 0)

	queue = append(queue, root)

	for len(queue) != 0 {
		nextQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		result = append(result, queue[len(queue)-1].Val)
		queue = nextQueue
	}
	return result
}
