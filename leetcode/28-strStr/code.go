package main

//使用类似进制的滚动哈希
func strStr(haystack string, needle string) int {
	//划定哈希值范围，超过则取余，为负则加上base
	const base = 10000000
	m, n := len(needle), len(haystack)
	if m > n {
		return -1
	}
	if m == 0 {
		return 0
	}

	power := 1
	for i := 1; i < m; i++ {
		power = (power * 31) % base
	}
	targetCode := hash(needle)
	firstCode := hash(haystack[0:m])
	if targetCode == firstCode {
		return 0
	}
	for i := 1; i < n-m+1; i++ {
		firstCode = firstCode - (int(haystack[i-1])*power)%base
		if firstCode < 0 {
			firstCode += base
		}
		firstCode = (firstCode*31 + int(haystack[i+m-1])) % base
		//double check

		if firstCode == targetCode {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}

//计算字符串hash
func hash(s string) int {
	base := 10000000
	code := 0
	for i, lt := 0, len(s); i < lt; i++ {
		code = (code*31 + int(s[i])) % base
	}
	return code
}

func main() {
	strStr("mississippi", "issi")
}
