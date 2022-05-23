package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0

提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/
/*
1. 排序
2. 取3个值：最左值，最右值，移动值
*/
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
