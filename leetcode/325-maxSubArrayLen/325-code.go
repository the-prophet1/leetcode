package main

import "fmt"

/*
给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。

注意:
 nums 数组的总和是一定在 32 位有符号整数范围之内的。

示例 1:

输入: nums = [1, -1, 5, -2, 3], k = 3
输出: 4
解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。
示例 2:

输入: nums = [-2, -1, 2, 1], k = 1
输出: 2
解释: 子数组 [-1, 2] 和等于 1，且长度最长。
进阶:
你能使时间复杂度在 O(n) 内完成此题吗?
*/

//暴力双重循环找最大长度O(n^2)
func maxSubArrayLen(nums []int, k int) int {
	var length = 0
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				if length < j-i+1 {
					length = j - i + 1
				}
			}
		}
	}
	return length
}

//利用前缀数组以及hash表，复杂度为O(3n)= O(n)
func maxSubArrayLen2(nums []int, k int) int {
	prefixArray := make([]int, 0)
	hash := make(map[int]int)
	length := 0
	sum := 0
	for i := 0; i < len(nums); i++ { //O(n)
		sum += nums[i]
		prefixArray = append(prefixArray, sum) //获取前缀数组和
	}

	for i := 0; i < len(prefixArray); i++ { //O(n)
		_, ok := hash[prefixArray[i]]
		if ok == false { //对于重复的前缀数组和，取先被录入的，因为是要求最大的k的长度
			hash[prefixArray[i]] = i
		}
	}
	for i := len(prefixArray) - 1; i >= 0; i-- { //O(n)
		if prefixArray[i] == k {
			return i + 1
		}
		index, ok := hash[prefixArray[i]+k]
		if ok == true {
			if i-index+1 > length {
				length = i - index + 1
			}
		}
	}
	return length
}

func main() {
	fmt.Println(maxSubArrayLen2([]int{1, -1, 5, -2, 3}, 3))
	fmt.Println(maxSubArrayLen2([]int{-2, -1, 2, 1}, 1))
	fmt.Println(maxSubArrayLen2([]int{0, 0, 1, 0, 1}, 2))
}
