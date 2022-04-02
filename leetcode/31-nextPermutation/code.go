package main

import (
	"fmt"
	"sort"
)

/*
整数数组的一个 排列就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,4,3,2] => [2,4,3,1] => [2,1,3,4]
输出：[2,1,3]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,2,3,4]
输出：[1,2,4,3]

输入：
[4,2,0,2,3,2,0]
输出：
[4,2,2,0,0,2,3]
预期结果：
[4,2,0,3,0,2,2]

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
思路
我们需要将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列。

同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。当交换完成后，「较大数」右边的数需要按照升序重新排列。
这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小。

*/

func nextPermutation(nums []int) {
	left := 0
	leftIndex := -1
	//从右往左寻找非降序的元素.即找一个较小的数尽可能靠右
	for index := len(nums) - 1; index >= 0; index-- {
		if left <= nums[index] {
			left = nums[index]
		} else {
			leftIndex = index
			break
		}
	}

	if leftIndex == -1 {
		sort.Ints(nums)
		return
	}

	//以nums[leftindex]为基点找到一个比nums[leftindex]大的数但尽可能小
	rightIndex := -1
	for index := len(nums) - 1; index >= 0; index-- {
		if nums[leftIndex] < nums[index] {
			rightIndex = index
			break
		}
	}
	nums[leftIndex], nums[rightIndex] = nums[rightIndex], nums[leftIndex]

	sort.Ints(nums[leftIndex+1:])
}

func main() {
	nums := []int{4, 2, 0, 2, 3, 2, 0}
	nextPermutation(nums)
	fmt.Println(nums)
}
