package main

import "fmt"

/**
53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。


示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

1 <= nums.length <= 105
-10^4 <= nums[i] <= 10^4

思路：使用动态规划求出前n的子数组的最大值：[-2,1,-2,4,3,5,6,1,5]
状态转义方程:
dp[i] 保存当前到该位置的最大值
if dp[i - 1] >= 0; dp[i] = dp[i - 1] + nums[n]
if dp[i - 1] < 0;  dp[i] = nums[n]
*/

func maxSubArray(nums []int) int {
	record := make([]int, len(nums))
	res := -100000
	record[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if record[i-1] < 0 {
			record[i] = nums[i]
		} else {
			record[i] = record[i-1] + nums[i]
		}
	}
	for i := 0; i < len(record); i++ {
		if record[i] > res {
			res = record[i]
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
