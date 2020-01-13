package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nextPtr := head

	for i := 0; i < n; i++ {
		nextPtr = nextPtr.Next
	}
	if nextPtr == nil {
		head = head.Next
		return head
	}
	prevPtr := head

	for nextPtr.Next != nil {
		nextPtr = nextPtr.Next
		prevPtr = prevPtr.Next
	}
	prevPtr.Next = prevPtr.Next.Next
	return head
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	head := removeNthFromEnd(node1, 5)
	fmt.Println(head)
	return
}
