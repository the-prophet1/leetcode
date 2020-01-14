package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	index := 0

	left := index + 1
	right := len(nums) - 1

	var minNum = math.MaxInt32

	for ; index < len(nums); index++ {
		left = index + 1
		right = len(nums) - 1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-minNum)) {
				minNum = sum
			}
			if sum > target { //sum太大了，所以right--可以减少sum的大小
				right--
			} else if sum < target { //sum太小了，所以left++可以增大sum的大小
				left++
			} else { //sum刚好等于target则直接返回就行了
				return minNum
			}
		}
	}
	return minNum
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
