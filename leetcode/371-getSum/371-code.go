package main

// a^b 为不进位的和
// a&b 计算出进位
// (a & b) << 1 把进位左移，赋值给b,继续用进位进行异或操作，求不进位的和

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
