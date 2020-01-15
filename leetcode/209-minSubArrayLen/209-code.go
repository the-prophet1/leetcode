package main

import (
	"fmt"
	"math"
)

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= s {
			return 1
		} else {
			return 0
		}
	}
	if nums[0] >= s {
		return 1
	}

	var minLength = math.MaxInt64
	var left = 0
	var right = 1
	sum := nums[left] + nums[right]
	for left <= right {
		if sum >= s {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		} else if sum < s {
			right++
			if right >= len(nums) {
				break
			}
			sum += nums[right]
		}
	}
	if minLength == math.MaxInt64 {
		return 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
