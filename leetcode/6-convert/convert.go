package main

import "fmt"

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。

比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"


提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000
*/

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}

	gap := (numRows - 1) * 2
	res := make([]byte, 0)
	for i := 0; i < numRows; i++ { // 枚举矩阵的行
		for j := 0; i+j < len(s); j += gap {
			res = append(res, s[i+j]) //当前周期第一个字符
			if 0 < i && i < numRows-1 && j+gap-i < len(s) {
				res = append(res, s[j+gap-i])
			}
		}
	}

	return string(res)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
