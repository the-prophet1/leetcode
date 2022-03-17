package main

import (
	"fmt"
	"math"
)

//数字反转使用求余的方式既能反转整个数字
func reverse(x int) int {

	var ret = 0

	for x != 0 {
		ret *= 10
		ret = ret + x%10
		x /= 10
	}
	if ret > math.MaxInt32 {
		return 0
	}
	if ret < math.MinInt32 {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(reverse(123456789))
}
