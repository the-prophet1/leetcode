package main

/*
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
*/

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		//确保 a的长度大于b
		return addBinary(b, a)
	}

	add1 := []byte(a)
	add2 := []byte(b)
	reverse(add1)
	reverse(add2)

	var carry byte = 0
	for i := 0; i < len(add2); i++ {
		sum := (add1[i] - '0') + (add2[i] - '0') + carry
		if sum >= 2 {
			add1[i] = sum - 2 + '0'
			carry = 1
		} else {
			add1[i] = sum + '0'
			carry = 0
		}
	}

	if carry == 1 {
		for i := len(add2); i < len(add1); i++ {
			sum := (add1[i] - '0') + carry
			if sum == 2 {
				add1[i] = '0'
				carry = 1
			} else {
				add1[i] = sum + '0'
				carry = 0
				break
			}
		}
	}

	if carry == 1 {
		add1 = append(add1, '1')
	}
	reverse(add1)
	return string(add1)
}

func reverse(nums []byte) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}

func main() {
	addBinary("101", "1")
}
