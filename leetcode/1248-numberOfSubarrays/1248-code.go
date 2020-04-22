package main

import "fmt"

/**
给你一个整数数组 nums 和一个整数 k。
如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
请返回这个数组中「优美子数组」的数目。
示例 1：

输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16


*/

func numberOfSubarrays(nums []int, k int) int {
	indexs := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			indexs = append(indexs, i)
		}
	}
	if len(indexs) < k {
		return 0
	}

	sum := 0
	for i := 0; i <= len(indexs)-k; i++ {
		left := indexs[i]
		right := indexs[i+k-1]
		leftCount := 0
		rightCount := 0
		if i == 0 {
			leftCount = left
		} else {
			leftCount = left - indexs[i-1] - 1
		}
		if i+k == len(indexs) {
			rightCount = len(nums) - indexs[i+k-1] - 1
		} else {
			rightCount = indexs[i+k] - right - 1
		}
		sum += (leftCount + 1) * (rightCount + 1)
	}
	return sum
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}
