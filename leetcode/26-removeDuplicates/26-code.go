package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count++
			continue
		} else {
			if nums[i] == nums[i-1] {
				continue
			} else {
				nums[count] = nums[i]
				count++
			}
		}
	}
	return count
}

func main() {
	data := []int{0, 0, 0, 1, 1, 2, 3, 4, 4, 5, 5, 6, 6, 8, 8, 9}
	l := removeDuplicates(data)
	fmt.Println(l, data)
}
