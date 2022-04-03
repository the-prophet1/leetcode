package main

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。


思路： 用每行出现的数字表、每列出现的数字表、每个九宫格出现的数字表进行匹配，当出现重复时，则为false，否则为true
*/

func isValidSudoku(board [][]byte) bool {
	//每行出现数字
	rows := [9][9]byte{}

	//每列出现的数字
	columns := [9][9]byte{}

	//每个九宫格出现的数字
	box := [9][9]byte{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' { //有数字
				if rows[i][board[i][j]-'1'] != 0 || columns[j][board[i][j]-'1'] != 0 || box[(i/3*3)+j/3][board[i][j]-'1'] != 0 {
					return false
				} else {
					rows[i][board[i][j]-'1'] = 1
					columns[j][board[i][j]-'1'] = 1
					box[(i/3*3)+j/3][board[i][j]-'1'] = 1
				}
			}
		}
	}
	return true
}

func main() {

}
