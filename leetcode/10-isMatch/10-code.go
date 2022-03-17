package main

/*
给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。

*/

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}

	//查看首元素是否一致
	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	//如果下一个字符是*
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}

	return firstMatch && isMatch(s[1:], p[1:])
}

func main() {
	isMatch("ab", ".*c")

}
