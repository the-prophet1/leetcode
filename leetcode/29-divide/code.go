package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	// 当 -2^31 / -1 会溢出，这是因为最大的int32是 2^31 - 1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := div(abs(dividend), abs(divisor))

	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		return result
	} else {
		return -result
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func div(dividend, divisor int) int {
	res := 0
	// 当 被除数 大于等于 除数
	for dividend >= divisor {
		value := divisor
		quotient := 1

		for dividend >= value {
			quotient += quotient
			value += value
		}
		quotient = quotient >> 1
		value = value >> 1
		res += quotient
		dividend -= value
	}

	return res
}

func main() {
	fmt.Println(divide(1100540749, -1090366779))
}
