package main

import "fmt"

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
				(m列 × n行)
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？

示例 1:
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右

示例 2:
输入: m = 7, n = 3
输出: 28

思路：
*/

func uniquePaths(m int, n int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		matrix[i][0] = 1
	}
	for i := 0; i < m; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}
	return matrix[n-1][m-1]
}

func main() {
	fmt.Println(uniquePaths(2, 3))
}
