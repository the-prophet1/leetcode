package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//----------------------------------------------------------------------------------------
//方法1:
//利用搜索二叉树的中序遍历是从小到大排列的
//中序遍历
func inorderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversal(root.Right, result)
}

func isValidBST(root *TreeNode) bool {
	var result []int
	inorderTraversal(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

//----------------------------------------------------------------------------------------
//方法2：
//利用搜索二叉树的左子树和右子树也为搜索二叉树，且左节点使用小于根节点，右节点始终大于根节点
//							下界	   上界
func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	if helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper) {
		return true
	}
	return false
}

func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func main() {

}
