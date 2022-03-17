package main

/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
*/

var numberMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return l([]string{}, digits)
}

func l(res []string, digits string) []string {
	if len(digits) == 0 {
		return res
	}
	chars := numberMap[digits[0]]
	if len(res) == 0 {
		res = chars
	} else {
		newRes := make([]string, 0)
		for _, char := range chars {
			for _, re := range res {
				newRes = append(newRes, re+char)
			}
		}
		res = newRes
	}
	return l(res, digits[1:])
}
