package main

import "fmt"

/*
给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。





输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/

func trap(height []int) int {
	leftMaxHeight := make([]int, len(height))
	rightMaxHeight := make([]int, len(height))
	maxHeight := 0
	//计算当前位置左边的墙的最高
	for i := 0; i < len(height); i++ {
		leftMaxHeight[i] = maxHeight
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	maxHeight = 0

	//计算当前位置右边的墙的最高
	for i := len(height) - 1; i >= 0; i-- {
		rightMaxHeight[i] = maxHeight
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}

	res := 0
	for i := 1; i < len(height)-1; i++ {
		//计算每个下标可以接多少水
		if height[i] < min(leftMaxHeight[i], rightMaxHeight[i]) {
			res += min(leftMaxHeight[i], rightMaxHeight[i]) - height[i]
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var height = []int{1}
	fmt.Println(trap(height))
}
