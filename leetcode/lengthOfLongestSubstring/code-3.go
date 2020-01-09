package main

//实现类似滑动窗口方式判断最长子串，通过下标指向滑动窗口的左边框和右边框
//每遇到一个字符则与之前边框内的字符进行比较
//若为新字符则把右边框右移
//若不为新字符则把左边框移到边框内相同字符的右边一位并将其设置为左边框
//时间复杂度为O(n),空间复杂度为O(1)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 { //字符串长度为１
		return 1
	}
	if len(s) == 0 { //字符串长度为0
		return 0
	}
	leftIndex := 0
	rightIndex := 1
	max := 1 //保存一个最长的长度

	for rightIndex < len(s) {
		current := 1
		for i := leftIndex; i < rightIndex; i++ {
			if s[i] == s[rightIndex] { //相同字符则将leftIndex设置为相同字符右边一位
				leftIndex = i + 1
				break
			}
			current++
			if current > max {
				max = current
			}
		}
		rightIndex++
	}
	return max
}
