package main

import "fmt"

/**
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

思路1：枚举的方式
思路2：使用前缀的方式，k的值为2个前缀和相减的区间，所以 k = pre[m] - pre[n]
*/

func subarraySum(nums []int, k int) int {
	count := 0
	var current int
	for i := 0; i < len(nums); i++ {
		current = nums[i]
		if current == k {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			current += nums[j]
			if current == k {
				count++
			}
		}
	}
	return count
}

func subarraySum1(nums []int, k int) int {
	pre := 0
	m := map[int]int{}
	m[0] = 1
	count := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		_, ok := m[pre-k]
		if ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum1([]int{1, 1, 1}, 2))
}
