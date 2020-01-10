package main

import "fmt"

//双指针法找面积
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var max = 0
	current := 0
	for left < right {
		if height[right] > height[left] {
			width := right - left
			current = width * height[left]
			if current > max {
				max = current
			}
			left++
		} else {
			width := right - left

			current = width * height[right]
			if current > max {
				max = current
			}
			right--
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
