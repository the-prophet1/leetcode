package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pt1 := l1
	pt2 := l2
	ret := &ListNode{}
	current := ret
	isCarry := false //进位符
	for pt1 != nil || pt2 != nil {
		if isCarry { //保存下次相加是否提前加１
			current.Val++
			isCarry = false
		}
		if pt1 != nil {
			current.Val += pt1.Val
			if current.Val >= 10 { //进行进位
				current.Val -= 10
				isCarry = true
			}
			pt1 = pt1.Next
		}
		if pt2 != nil {
			current.Val += pt2.Val
			if current.Val >= 10 { //进行进位
				current.Val -= 10
				isCarry = true
			}
			pt2 = pt2.Next
		}
		if pt1 == nil && pt2 == nil {
			if isCarry { //最后一次进位
				current.Next = &ListNode{1, nil}
			}
			break
		}
		current.Next = &ListNode{}
		current = current.Next
	}
	return ret
}

func main() {
	l3 := &ListNode{3, nil}
	l2 := &ListNode{4, l3}
	l1 := &ListNode{2, l2}

	l33 := &ListNode{4, nil}
	l22 := &ListNode{6, l33}
	l11 := &ListNode{5, l22}

	addTwoNumbers(l1, l11)
}
