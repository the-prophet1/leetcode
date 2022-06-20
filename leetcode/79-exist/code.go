package main

import "fmt"

/*
79. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false


提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成


进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

思路：使用DFS进行搜索，并用一个相同的数组记录已经走过的路径
*/

func exist(board [][]byte, word string) bool {
	wordByte := []byte(word)
	tags := make([][]bool, len(board))
	for i := 0; i < len(tags); i++ {
		tags[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if trySearch(board, tags, wordByte, i, j) == true {
				return true
			}
		}
	}
	return false
}

func trySearch(board [][]byte, tags [][]bool, wordByte []byte, row, column int) bool {
	if len(wordByte) == 0 {
		return true
	}
	if row < 0 || column < 0 {
		return false
	}
	if row > len(board)-1 || column > len(board[0])-1 {
		return false
	}
	if tags[row][column] == true {
		return false
	}
	if board[row][column] != wordByte[0] {
		return false
	}

	tags[row][column] = true

	// 尝试向上查询下一个单词
	res := trySearch(board, tags, wordByte[1:], row-1, column) ||
		// 尝试向下查询下一个单词
		trySearch(board, tags, wordByte[1:], row+1, column) ||
		// 尝试向左查询下一个单词
		trySearch(board, tags, wordByte[1:], row, column-1) ||
		// 尝试向右查询下一个单词
		trySearch(board, tags, wordByte[1:], row, column+1)
	if res == false {
		tags[row][column] = false
	}
	return res
}

func main() {

	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	fmt.Println(exist(board, word))
}
