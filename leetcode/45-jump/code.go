package main

import "fmt"

/*
45. 跳跃游戏 II
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2
[2,3,0,1,4]
[2,4,2,4,8]

提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	jumpArray := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		jumpArray[i] = i + nums[i]
	}

	count := 1
	length := len(nums) - 1
	target := jumpArray[0]
	index := 0
	for {
		if target >= length {
			return count
		}
		current := 0
		current, target = max(jumpArray[index+1 : target+1])
		index += 1 + current
		count++
	}
}

func max(sli []int) (int, int) {
	index := 0
	current := -1
	for i, num := range sli {
		if num > current {
			current = num
			index = i
		}
	}
	return index, current
}

func main() {
	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(jump(nums))
}
