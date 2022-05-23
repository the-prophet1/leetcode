package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回[-1, -1]。

进阶：

你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109

*/

/*
 [1,2,2,2,3]
*/

func median(a, b int) int {
	return (a + b) / 2
}

func binaryRightSearch(nums []int, x int) int {
	left := 0
	right := len(nums) - 1

	ret := -1

	for left <= right {
		if nums[median(left, right)] == x {
			//此处将二分查找改为对右侧的数组继续进行查找
			ret = median(left, right)
			left = median(left, right) + 1
		} else if nums[median(left, right)] > x {
			right = median(left, right) - 1
		} else if nums[median(left, right)] < x {
			left = median(left, right) + 1
		}
	}
	return ret
}

func binaryLeftSearch(nums []int, x int) int {
	left := 0
	right := len(nums) - 1

	ret := -1

	for left <= right {
		if nums[median(left, right)] == x {
			//此处将二分查找改为对左侧的数组继续进行查找
			ret = median(left, right)
			right = median(left, right) - 1
		} else if nums[median(left, right)] > x {
			right = median(left, right) - 1
		} else if nums[median(left, right)] < x {
			left = median(left, right) + 1
		}
	}
	return ret
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	res[0] = binaryLeftSearch(nums, target)
	res[1] = binaryRightSearch(nums, target)
	return res
}

func main() {
	nums := []int{1, 2, 2, 2, 3}
	fmt.Println(searchRange(nums, 2))
}
