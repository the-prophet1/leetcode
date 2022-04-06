package main

import "fmt"

/*
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

思路：
以'回'型的方式打印数组
*/

func spiralOrder(matrix [][]int) []int {
	topIndex := 0
	bottomIndex := len(matrix) - 1
	leftIndex := 0
	rightIndex := len(matrix[0]) - 1
	res := make([]int, 0)
	for bottomIndex > topIndex && rightIndex > leftIndex {
		//保存上行
		for i := leftIndex; i < rightIndex; i++ {
			res = append(res, matrix[topIndex][i])
		}

		//保存右列
		for i := topIndex; i < bottomIndex; i++ {
			res = append(res, matrix[i][rightIndex])
		}

		//保存下行
		for i := rightIndex; i > leftIndex; i-- {
			res = append(res, matrix[bottomIndex][i])
		}

		// 保存左行
		for i := bottomIndex; i > topIndex; i-- {
			res = append(res, matrix[i][leftIndex])
		}
		topIndex++
		bottomIndex--
		leftIndex++
		rightIndex--
	}

	if topIndex == bottomIndex {
		for i := leftIndex; i <= rightIndex; i++ {
			res = append(res, matrix[topIndex][i])
		}
	} else if leftIndex == rightIndex {
		for i := topIndex; i <= bottomIndex; i++ {
			res = append(res, matrix[i][leftIndex])
		}
	}

	return res
}

func main() {
	nums := [][]int{{1}}
	fmt.Println(spiralOrder(nums))
}
