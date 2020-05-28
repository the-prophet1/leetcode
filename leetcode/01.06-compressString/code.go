package main

import (
	"fmt"
	"strconv"
)

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。

*/

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	var s []byte

	var currentChar uint8 = 0
	cnt := 0
	for i := 0; i < len(S); i++ {
		if currentChar == 0 {
			currentChar = S[i]
			cnt++
		} else if currentChar == S[i] {
			cnt++
		} else {
			s = append(s, S[i-1])
			s = append(s, []byte(strconv.Itoa(cnt))...)
			cnt = 1
			currentChar = S[i]
		}
	}
	s = append(s, S[len(S)-1])
	s = append(s, []byte(strconv.Itoa(cnt))...)

	if len(s) > len(S) {
		return S
	} else {
		return string(s)
	}
}

func main() {
	fmt.Println(compressString("bb"))
}
