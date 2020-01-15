package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法1：
//递归的方式进行2个树之间的比较
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {

}
