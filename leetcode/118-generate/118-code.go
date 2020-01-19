package main

import "fmt"

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
nums:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

思路：
1.上述的第三层的第二个数字:nums[2][1],是有第二层的1个数字和第2个数字组成，即:nums[2][1] = nums[0][0] + nums[0][1]
2.第四层的第二个数字:nums[3][1] = nums[2][0] + nums[2][1]
3.第四层的第3个数字: nums[3][2] = nums[2][1] + nums[2][2]
所以中间非1的数字都是由 nums[n][m] = nums[n-1][m-1] +nums[n-1][m],其中n >= 2 ,m > 1 && m < n
*/

func generate(numRows int) [][]int {
	YangHuiTriangle := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		nums := make([]int, 0)
		for j := 0; j <= i; j++ {
			nums = append(nums, 1)
		}
		YangHuiTriangle = append(YangHuiTriangle, nums)
	}

	if numRows <= 2 {
		return YangHuiTriangle
	}

	for n := 2; n < numRows; n++ {
		for m := 1; m < n; m++ {
			YangHuiTriangle[n][m] = YangHuiTriangle[n-1][m-1] + YangHuiTriangle[n-1][m]
		}
	}

	return YangHuiTriangle
}

func main() {
	fmt.Println(generate(5))
}
