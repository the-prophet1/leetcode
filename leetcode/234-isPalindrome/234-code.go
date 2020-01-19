package main

import "fmt"

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fastPoint := head
	slowPoint := head
	var ptr, p *ListNode
	for fastPoint != nil && fastPoint.Next != nil {
		ptr = slowPoint
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next

		ptr.Next = p
		p = ptr
	}

	if fastPoint != nil { //链表长度为偶数
		slowPoint = slowPoint.Next
	}
	for p != nil {
		if slowPoint.Val != p.Val {
			return false
		}
		p = p.Next
		slowPoint = slowPoint.Next
	}
	return true
}

func main() {
	node6 := &ListNode{1, nil}
	node5 := &ListNode{2, node6}
	node4 := &ListNode{3, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	fmt.Println(isPalindrome(node1))
}
