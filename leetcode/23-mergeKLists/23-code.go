package main

import "fmt"

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
思路1：循环合并每个链表
思路2：使用归并的方式两两合并最后得到一个链表
思路3：为每个链表维护的头部一个最小堆
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	//思路1
	//var res *ListNode
	//for i := 0; i < len(lists); i++ {
	//	res = mergeList(res,lists[i])
	//}

	//思路2
	var mlists []*ListNode
	var left = 0
	var right = len(lists) - 1
	for left < right {
		list := mergeList(lists[left], lists[right])
		mlists = append(mlists, list)
		left++
		right--
	}
	if left == right {
		mlists = append(mlists, lists[left])
	}
	return mergeKLists(mlists)
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var preHead = &ListNode{0, nil}
	var current = preHead
	var p1 = list1
	var p2 = list2

	for p1 != nil && p2 != nil {
		var p *ListNode
		if p1.Val > p2.Val {
			p = p2
			p2 = p2.Next
			current.Next = p
			current = p
		} else {
			p = p1
			p1 = p1.Next
			current.Next = p
			current = p
		}
	}
	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}
	return preHead.Next
}

func main() {
	node3 := &ListNode{5, nil}
	node2 := &ListNode{4, node3}
	node1 := &ListNode{1, node2}

	nodez := &ListNode{4, nil}
	nodey := &ListNode{3, nodez}
	nodex := &ListNode{1, nodey}

	nodeZ := &ListNode{3, nil}
	nodeX := &ListNode{1, nodeZ}

	var lists []*ListNode
	lists = append(lists, node1)
	lists = append(lists, nodex)
	lists = append(lists, nodeX)

	res := mergeKLists(lists)
	p := res
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
