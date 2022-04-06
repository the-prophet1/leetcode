package main

/*
55. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 10^4
0 <= nums[i] <= 10^5

思路：
使用record 保存每个下标所能到达的最远位置，
遍历record，当下标超过record记录的最长长度时，说明已经不可达了，返回false
*/

func canJump(nums []int) bool {
	record := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		record[i] = nums[i] + i
	}
	m := record[0]
	for i := 1; i < len(record); i++ {
		if i > m {
			return false
		}
		m = max(m, record[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
