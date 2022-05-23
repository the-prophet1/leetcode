package main

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

// 将值进行反转后存入
func reverseKGroup(head *ListNode, k int) *ListNode {
	data := make([]int, k)
	ptr := head
	nextPtr := head
	for nextPtr != nil {
		for i := 0; i < k; i++ {
			if nextPtr == nil {
				return head
			}
			data[i] = nextPtr.Val
			nextPtr = nextPtr.Next
		}
		reverseData(data)
		for i := 0; i < k; i++ {
			ptr.Val = data[i]
			ptr = ptr.Next
		}
	}
	return head
}

func reverseData(data []int) {
	i := 0
	j := len(data) - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	reverseKGroup(node1, 2)
}
