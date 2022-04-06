package main

/*
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：


输入：head = [0,1,2], k = 4
输出：[2,0,1]


提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	vals := make([]int, 0)
	ptr := head
	for ptr != nil {
		vals = append(vals, ptr.Val)
		ptr.Next = ptr
	}
	k = k % len(vals)

	start := len(vals) - k

	ptr = head
	for i := start; i < len(vals); i++ {
		ptr.Val = vals[i]
		ptr = ptr.Next
	}
	i := 0
	for ptr != nil {
		ptr.Val = vals[i]
		ptr = ptr.Next
	}

	return head
}

func main() {

}
