package main

import "fmt"

/**
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

解析：构造一个记录该位置左上部分的正方形长度的矩阵
			 0 0 0 0 0 0
1 0 1 0 0    0 1 0 1 0 0
1 0 1 1 1    0 1 0 1 1 1
1 1 1 1 1    0 1 1 1 2 2
1 0 1 1 1    0 1 0 1 2 3
*/

func isSquare(matrix [][]byte, x int, y int) bool {
	if matrix[x-1][y] >= matrix[x-1][y-1] &&
		matrix[x][y-1] >= matrix[x-1][y-1] {
		return true
	}
	return false
}

func maximalSquare(matrix [][]byte) int {
	buf := make([][]byte, len(matrix)+1)
	buf[0] = make([]byte, len(matrix[0])+1)

	for index, sl := range buf {
		if index == 0 {
			continue
		}
		sl = append(sl, 0)
		sl = append(sl, matrix[index-1]...)
		buf[index] = sl
	}

	maxsqlen := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				if buf[i][j] == 0 {
					if maxsqlen < 1 {
						maxsqlen = 1
					}
				} else {
					flag := isSquare(buf, i+1, j+1)
					if flag == true {
						buf[i+1][j+1] = buf[i][j] + 1
						if int(buf[i][j]+1) > maxsqlen {
							maxsqlen = int(buf[i][j] + 1)
						}
					}
				}
			}
		}
	}
	return maxsqlen * maxsqlen
}

func main() {
	matrix := [][]byte{{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0}}
	fmt.Println(maximalSquare(matrix))
}
