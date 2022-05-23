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
			//当右边高度高于左边高度时,计算以下当前包含的面积
			width := right - left
			current = width * height[left]
			if current > max {
				max = current
			}
			//并且由于右边高于左边，所以只可能当左边高于右边的时候。才可能包含更多的面积，所以left++
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
