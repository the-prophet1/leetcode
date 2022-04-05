package main

func solveSudoku(board [][]byte) {
	//每行出现数字
	rows := [9][9]byte{}

	//每列出现的数字
	columns := [9][9]byte{}

	//每个九宫格出现的数字
	box := [9][9]byte{}
	filledIn := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' { //有数字
				rows[i][board[i][j]-'1'] = 1
				columns[j][board[i][j]-'1'] = 1
				box[(i/3*3)+j/3][board[i][j]-'1'] = 1
				filledIn++
			}
		}
	}

	for filledIn != 81 {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == '.' {
					//不是数字则尝试填写
					checks := [9]byte{}
					for k := 0; k < 9; k++ {
						checks[k] = rows[i][k] | columns[j][k] | box[(i/3*3)+j/3][k]
					}
					var maybe byte = 0
					count := 0 //填写次数
					for k := 0; k < 9; k++ {
						if checks[k] == 0 {
							maybe = byte(k)
							count++
							if count > 1 {
								break
							}
						}
					}
					if count == 1 {
						board[i][j] = '1' + maybe
						rows[i][board[i][j]-'1'] = 1
						columns[j][board[i][j]-'1'] = 1
						box[(i/3*3)+j/3][board[i][j]-'1'] = 1
						filledIn++
					}
				}
			}
		}
	}

}

func main() {
	sudu := [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}}

	solveSudoku(sudu)
	return
}
