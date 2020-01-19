package main

import (
	"fmt"
	"math"
)

func getMin(strs []string) (min int) {
	min = math.MaxInt64
	for i := 0; i < len(strs); i++ {
		if min > len(strs[i]) {
			min = len(strs[i])
		}
	}
	return
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	min := getMin(strs)
	var res = make([]byte, 0)
	for i := 0; i < min; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if c != strs[j][i] {
				return string(res)
			}
		}
		res = append(res, c)
	}
	return string(res)
}

func main() {
	s := []string{}
	fmt.Println(longestCommonPrefix(s))
}
