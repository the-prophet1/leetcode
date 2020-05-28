package main

import "math"

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

思路：使用动态规划求出前n的子数组的最大值：[-2,1,-2,4,3,5,6,1,5]
*/

func maxSubArray(nums []int) int {
	var recordArray []int
	var result = math.MinInt64
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			recordArray = append(recordArray, nums[i])
		} else {
			if nums[i]+recordArray[i-1] > nums[i] {
				recordArray = append(recordArray, nums[i]+recordArray[i-1])
			} else {
				recordArray = append(recordArray, nums[i])
			}
		}
	}

	for i := 0; i < len(recordArray); i++ {
		if recordArray[i] > result {
			result = recordArray[i]
		}
	}
	return result
}
