package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var max = 1 //最大子串长度
	var maxToIndex float64
	var result []byte
	//单数个数的最大回文串

	for index := 0; index < len(s); index++ {
		left := index - 1
		right := index + 1
		current := 1
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				current += 2
				left--
				right++
				if current > max {
					max = current
					maxToIndex = float64(index)
				}
			} else {
				break
			}

		}
	}

	//双数个数的最大回文串
	for index := 0.5; index < float64(len(s)); index++ {
		left := int(index - 0.5)
		right := int(index + 0.5)
		current := 0
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				current += 2
				left--
				right++
				if current > max {
					max = current
					maxToIndex = index
				}
			} else {
				break
			}

		}
	}

	if int(maxToIndex*2)%2 == 0 { //子串为单数个
		offset := max / 2
		index := int(maxToIndex) - offset
		for i := 0; i < max; i++ {
			result = append(result, s[index+i])
		}

	} else { //子串为偶数个
		index := int(maxToIndex + 0.5)
		offset := max / 2
		index = index - offset
		for i := 0; i < max; i++ {
			result = append(result, s[index+i])
		}
	}

	return string(result)
}

func main() {
	fmt.Println(longestPalindrome("abcbc"))
}
