package main

/*
82. 删除排序链表中的重复元素 II
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
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

// 1 1 1
func deleteDuplicates(head *ListNode) *ListNode {
	flag := false
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head = head.Next
		} else {
			flag = true
		}
	}
	if !flag {
		return nil
	}
	ptr := head

	return head
}

func main() {
	node5 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 2, Next: &node5}
	node3 := ListNode{Val: 2, Next: &node4}
	node2 := ListNode{Val: 1, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	deleteDuplicates(&node1)
}
