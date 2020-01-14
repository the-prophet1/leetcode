package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归解法
func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsRecursive(next.Next)
	next.Next = head
	return next
}

//非递归解法
func swapPairs(head *ListNode) *ListNode {
	ptr := head
	if ptr == nil { //链表无节点
		return head
	}

	nextPtr := head.Next
	if nextPtr == nil { //链表只有一个节点
		return head
	}
	firstHead := nextPtr
	prevNode := ptr
	//链表有２个以上的节点
	for ptr != nil && nextPtr != nil {
		//交换２个节点的位置
		temp := nextPtr.Next
		nextPtr.Next = ptr
		ptr.Next = temp

		//修正ptr和nextPtr的位置
		temp = ptr
		ptr = nextPtr
		nextPtr = temp

		//将ptr和nextPtr指向后２个位置
		ptr = ptr.Next.Next
		if ptr == nil { //判断ptr是否到结尾了
			continue
		}
		nextPtr = nextPtr.Next.Next

		if nextPtr != nil { //判断nextPtr是否到结尾了
			prevNode.Next = nextPtr
			prevNode = ptr
		}

	}
	return firstHead
}

func main() {
	node6 := &ListNode{6, nil}
	node5 := &ListNode{5, node6}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	head := swapPairs(node1)
	head = swapPairsRecursive(head)
	fmt.Println(head)
}
