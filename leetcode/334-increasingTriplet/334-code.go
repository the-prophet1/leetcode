package main

import (
	"fmt"
	"math"
)

/**
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
*/

func increasingTriplet(nums []int) bool {
	var first int = math.MaxInt64
	var second int = math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < first {
			first = nums[i]
		} else if nums[i] < second && first < second {
			second = nums[i]
		} else if nums[i] > second {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{2, 3, 1, 2, 1}
	fmt.Println(increasingTriplet(arr))
}
