package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	var mergeList *ListNode = head

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			mergeList.Next = l2
			l2 = l2.Next
			mergeList = mergeList.Next
		} else {
			mergeList.Next = l1
			l1 = l1.Next
			mergeList = mergeList.Next
		}
	}

	if l1 == nil {
		mergeList.Next = l2
	} else {
		mergeList.Next = l1
	}
	return head.Next
}

func main() {
	lnode3 := &ListNode{4, nil}
	lnode2 := &ListNode{3, lnode3}
	lnode1 := &ListNode{1, lnode2}

	rnode2 := &ListNode{4, nil}
	rnode1 := &ListNode{2, rnode2}
	head := mergeTwoLists(lnode1, rnode1)
	fmt.Println(head)
}
