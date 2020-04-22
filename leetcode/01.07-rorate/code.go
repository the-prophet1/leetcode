package main

import "fmt"

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
]

方法1：把矩阵以圈的方式进行旋转

方法2：利用数学的原理，将矩阵对角翻转，再左右翻转
*/

func rotate(matrix [][]int) {
	sideLength := len(matrix)
	for i := 0; i < sideLength; i++ {
		//对角线翻转
		for j := i; j < sideLength; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		//左右翻转
		for k := 0; k < sideLength/2; k++ {
			matrix[i][k], matrix[i][sideLength-k-1] = matrix[i][sideLength-k-1], matrix[i][k]
		}
		fmt.Println(matrix[i])
	}
}

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
}
