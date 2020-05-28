package main

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:

输入:[2,3,-2,1]

输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

*/

func Min(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func maxProduct(nums []int) int {
	imax := nums[0]
	imin := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax //此处转换当前的最小值和最大值，因为最大值乘上符号，会变成最小值，而最小值乘上负号会变成最小值
		}
		imax = Max(imax*nums[i], nums[i])
		imin = Min(imin*nums[i], nums[i])

		max = Max(max, imax)
	}
	return max
}

func main() {
	maxProduct([]int{2, 3, -2, 1})
}
