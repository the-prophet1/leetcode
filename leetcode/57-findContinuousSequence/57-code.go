package main

import "fmt"

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
*/

func findContinuousSequence(target int) [][]int {
	limit := target / 2
	sum := 0
	res := make([][]int, 0, 0)
	for i := 1; i <= limit+1; i++ {
		arr := make([]int, 0, 0)
		for j := i; j <= limit+1; j++ {
			sum += j
			arr = append(arr, j)
			if sum > target {
				sum = 0
				break
			}
			if sum == target {
				res = append(res, arr)
				sum = 0
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
