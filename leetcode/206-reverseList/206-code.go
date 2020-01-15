package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代方式完成链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head
	nextPtr := ptr.Next
	ptr.Next = nil
	for nextPtr != nil {
		temp := nextPtr.Next
		nextPtr.Next = ptr
		ptr = nextPtr
		nextPtr = temp
	}
	return ptr
}

//递归方式完成链表反转
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	Head := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return Head
}

func main() {

	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	ptr := reverseList(node1)
	fmt.Println(ptr)
}
