package main

import "sort"

/*
给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]
（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。



示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

提示：

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

*/

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	if len(nums) <= 3 {
		return res
	}

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			left := j + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left+1 < len(nums) && nums[left] == nums[left+1] {
						left++
					}
					left++
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				}
			}
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

func main() {
	var arr = []int{1, 0, -1, 0, -2, 2}
	fourSum(arr, 0)

}
