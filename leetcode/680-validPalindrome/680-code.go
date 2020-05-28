package main

/**
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "adfca"
输出: True
解释: 你可以删除c字符

注意：字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			l1 := i + 1
			r1 := j
			flag1 := true
			for l1 < r1 {
				if s[l1] == s[r1] {
					l1++
					r1--
				} else {
					flag1 = false
					break
				}
			}
			l2 := i
			r2 := j - 1
			flag2 := true
			for l2 < r2 {
				if s[l2] == s[r2] {
					l2++
					r2--
				} else {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func main() {
	validPalindrome("abc")
}
