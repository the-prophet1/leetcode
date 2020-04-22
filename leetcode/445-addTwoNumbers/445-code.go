package main

import "fmt"

/**
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

示例：
输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

思路1：对2个链表进行翻转，然后对2个链表的每个节点相加，得到新链表，再将新链表再次翻转后返回
思路2：使用栈将2个链表的数字进行压栈，完成压栈后再弹栈进行对2数进行相加，
	   压入第3个栈，完成压栈后再对第三个栈进行弹栈建立链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//思路2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0, 0)
	stack2 := make([]int, 0, 0)
	queue := make([]int, 0, 0)
	ptr1 := l1
	ptr2 := l2

	//对l1压栈
	for ptr1 != nil {
		stack1 = append(stack1, ptr1.Val)
		ptr1 = ptr1.Next
	}
	//对l2压栈
	for ptr2 != nil {
		stack2 = append(stack2, ptr2.Val)
		ptr2 = ptr2.Next
	}

	carry := 0
	nums1 := 0
	nums2 := 0
	//对queue压栈
	for len(stack1) != 0 || len(stack2) != 0 || carry != 0 {
		//对stack1弹栈
		if len(stack1) != 0 {
			nums1 = stack1[len(stack1)-1]
			stack1 = append(stack1[:len(stack1)-1])
		}
		//对stack2弹栈
		if len(stack2) != 0 {
			nums2 = stack2[len(stack2)-1]
			stack2 = append(stack2[:len(stack2)-1])
		}
		queue = append(queue, (nums1+nums2+carry)%10)
		//判断是否进位
		if nums1+nums2+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		nums1 = 0
		nums2 = 0
	}
	var ptr *ListNode = nil
	for len(queue) != 0 {
		nums := queue[0]
		queue = append(queue[1:len(queue)])
		node := &ListNode{nums, nil}
		node.Next = ptr
		ptr = node
	}
	return ptr
}

func main() {
	node4 := &ListNode{3, nil}
	node3 := &ListNode{4, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{7, node2}

	nodeZ := &ListNode{4, nil}
	nodeY := &ListNode{6, nodeZ}
	nodeX := &ListNode{5, nodeY}

	res := addTwoNumbers(node1, nodeX)
	fmt.Println(res)
}
