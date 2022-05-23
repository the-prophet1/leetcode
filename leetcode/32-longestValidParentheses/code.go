package main

import "fmt"

/*
给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

输入：s := "()(()()(("
输出：2


提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

使用一个栈进行匹配
栈中保存左括号的下标
*/

type stack struct {
	s []int
}

func longestValidParentheses(s string) int {
	st := newStack()
	res := 0
	tags := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			st.push(i)
		}
		if char == ')' {
			if !st.isNil() {
				//括号匹配上了
				tags[i] = true
				e := st.pop()
				tags[e] = true
			}
		}
	}
	current := 0
	for i := 0; i < len(tags); i++ {
		if tags[i] == true {
			current++
			if current > res {
				res = current
			}
		} else {
			current = 0
		}
	}
	return res
}

func newStack() *stack {
	return &stack{s: []int{}}
}

func (st *stack) push(index int) {
	st.s = append(st.s, index)
}

func (st *stack) isNil() bool {
	if len(st.s) == 0 {
		return true
	}
	return false
}

func (st *stack) pop() int {
	ret := st.s[len(st.s)-1]
	st.s = st.s[:len(st.s)-1]
	return ret
}

func main() {
	s := "()(()())(("
	fmt.Println(longestValidParentheses(s))

}
