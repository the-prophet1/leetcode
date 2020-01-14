package main

import "fmt"

func removeElement(nums []int, val int) int {
	var count = 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	var nums = []int{1, 2, 1, 2, 3, 4, 1, 2, 4, 3, 3, 2}
	size := removeElement(nums, 1)
	fmt.Println(nums, size)
}
