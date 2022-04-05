package main

import (
	"fmt"
	"strconv"
)

/*
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例2:

输入: num1 = "123", num2 = "456"
输出: "56088"


提示：

1 <= num1.length, num2.length <= 200
num1和 num2只能由数字组成。
num1和 num2都不包含任何前导零，除了数字0本身。

*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := ""
	for i := len(num2) - 1; i >= 0; i-- {
		multiplier := []byte(num1)
		carry := 0

		for j := len(multiplier) - 1; j >= 0; j-- {
			data := int(multiplier[j]-'0')*int(num2[i]-'0') + carry
			carry = data / 10      //进位数
			remainder := data % 10 //余数
			multiplier[j] = uint8(remainder) + '0'
		}
		addend := string(multiplier)
		if carry != 0 {
			addend = strconv.Itoa(carry) + addend
		}

		for j := len(num2) - 1; j > i; j-- {
			addend += "0"
		}
		res = addStrings(res, addend)
	}

	return res
}

func addStrings(num1 string, num2 string) string {
	//确保 len(num1) >= len(num2)
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}
	b1 := swap(num1)
	b2 := swap(num2)

	carry := 0

	for i := 0; i < len(num2); i++ {
		data := int(b1[i]-'0') + int(b2[i]-'0') + carry
		carry = data / 10      //进位数
		remainder := data % 10 //余数
		b1[i] = uint8(remainder) + '0'
	}

	if carry != 0 {
		for i := len(b2); i < len(b1); i++ {
			data := int(b1[i]-'0') + carry
			carry = data / 10
			remainder := data % 10 //余数
			b1[i] = uint8(remainder) + '0'
		}
	}
	if carry != 0 {
		b1 = append(b1, '1')
	}

	return string(swap(string(b1)))
}

func swap(s string) []byte {
	bytes := []byte(s)
	i := 0
	j := len(bytes) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return bytes
}

func main() {

	fmt.Println(multiply("10", "121"))
}
