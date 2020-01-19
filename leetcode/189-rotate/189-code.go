package main

import "fmt"

func rotate(nums []int, k int) {
	if len(nums) < k {
		k %= len(nums)
	}
	data := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, data)
}

func main() {
	res := []int{1, 2, 3, 4, 5, 6}
	rotate(res, 7)
	fmt.Println(res)
}
