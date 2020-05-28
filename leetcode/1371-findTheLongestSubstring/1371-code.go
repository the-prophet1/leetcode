package main

/**
给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，
即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。

示例 1：
s = "euleetminicoworoep"
输出：13
解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。

*/

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findTheLongestSubstring(s string) int {
	res, status := 0, 0
	pos := make([]int, 1<<5)
	for i := 1; i < len(pos); i++ {
		pos[i] = -1
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] >= 0 { //出现重复状态时
			res = Max(res, i+1-pos[status])
		} else {
			pos[status] = i + 1 //记录新的所有元音的状态出现的第一个位置
		}
	}
	return res
}

func main() {
	findTheLongestSubstring("elee")
}
