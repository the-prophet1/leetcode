package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	leftIndex := left - 1
	rightIndex := right - 1

	buf := make([]int, 0)
	ptr := head
	for ptr != nil {
		buf = append(buf, ptr.Val)
		ptr = ptr.Next
	}

	reverse(buf[leftIndex : rightIndex+1])
	ptr = head
	i := 0
	for ptr != nil {
		ptr.Val = buf[i]
		i++
	}
	return head
}

func reverse(data []int) {
	i := 0
	j := len(data) - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main() {

}
