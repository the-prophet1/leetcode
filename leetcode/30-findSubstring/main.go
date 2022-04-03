package main

import "fmt"

/*
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。



示例 1：

输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：
1,2,3
2,1,3
3,2,1
1,3,2
3,1,2
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
示例 3：

输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]

*/

func findSubstring1(s string, words []string) []int {
	res := make([]int, 0)
	length := len(words) * len(words[0])
	if length > len(s) {
		return res
	}

	var m = make(map[string]bool)
	var generateHash func([]string, string)
	//将words的全排列保存在hash中
	generateHash = func(s []string, exist string) {
		if len(s) == 0 {
			m[exist] = true
			return
		}

		for i := 0; i < len(s); i++ {

			//取出一个元素
			newexist := exist + s[i]
			//将剩余元素继续拼接
			st := make([]string, 0)
			st = append(st, s[:i]...)
			st = append(st, s[i+1:]...)
			generateHash(st, newexist)
		}

	}
	generateHash(words, "")

	//使用滑动窗口对字符串进行检查
	for i := 0; i <= len(s)-length; i++ {
		if m[s[i:i+length]] == true {
			res = append(res, i)
		}
	}
	return res
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	length := len(words) * len(words[0])
	wordLength := len(words[0])
	if length > len(s) {
		return res
	}

	//保存所有单词出现的次数
	var m = make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}

	//使用滑动窗口对字符串每个单词进行检查
	for i := 0; i <= len(s)-length; i++ {
		var mt = make(map[string]int)
		st := s[i : i+length]
		for j := 0; j < len(words); j++ {
			mt[st[j*wordLength:j*wordLength+wordLength]] += 1
		}
		if checkMap(m, mt) {
			res = append(res, i)
		}
	}
	return res
}

func checkMap(m1, m2 map[string]int) bool {
	for k, _ := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

func main() {
	s := []string{"foo", "bar"}
	fmt.Println(findSubstring("barfoothefoobar", s))
	return
}
