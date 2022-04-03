package main

import "fmt"

/**
实现pow(x, n)，即计算 x 的 n 次幂函数。

输入: 2.00000, 10
输出: 1024.00000

输入: 2.10000, 3
输出: 9.26100


说明:
-100.0 <x< 100.0
n是 32 位有符号整数，其数值范围是[−2^31,2^31 − 1] 。


思路1：暴力破解循环乘上n个x进行求解，为负号则乘上1/n个x进行求解
思路2：使用快速幂的方式进行求解
*/

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}

func main() {
	fmt.Println(myPow(2, 10))
}
