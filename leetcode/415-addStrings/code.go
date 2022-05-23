package main

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
	addStrings("456", "77")
}
