package main

/*
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

比较容易懂的做法是反转数组，加1，再反转
*/

func plusOne(digits []int) []int {
	reverse(digits)
	carry := 0
	for i := 0; i < len(digits); i++ {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
			break
		}
	}
	if carry == 1 {
		digits = append(digits, 1)
	}
	reverse(digits)
	return digits
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {

}
