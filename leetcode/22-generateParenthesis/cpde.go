package main

import "fmt"

/*



 */

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var dfs func(left int, right int, n int, s string)

	dfs = func(left int, right int, n int, s string) {
		if left == n && right == n {
			res = append(res, s)
			return
		}
		if left < n {
			dfs(left+1, right, n, s+"(")
		}
		if left > right { //只有左括号比右括号多，才能放入右括号
			dfs(left, right+1, n, s+")")
		}
	}

	dfs(0, 0, n, "")
	return res
}

func main() {

	fmt.Println(generateParenthesis(3))
}
