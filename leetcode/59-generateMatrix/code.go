package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	count := 1
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	i := 0
	foreach := n - 1
	for left < right {
		for j := i; j < foreach; j++ {
			matrix[top][j] = count
			count++
		}

		for j := i; j < foreach; j++ {
			matrix[j][right] = count
			count++
		}
		for j := i; j < foreach; j++ {
			matrix[bottom][n-j-1] = count
			count++
		}
		for j := i; j < foreach; j++ {
			matrix[n-j-1][left] = count
			count++
		}
		top++
		bottom--
		left++
		right--
		i++
		foreach--
	}

	if left == right {
		matrix[n/2][n/2] = count
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(2))

}
