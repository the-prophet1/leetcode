package main

/*
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

/*
方法1:
使用回溯法:
将例子nums[1,2,3]
转换为决策树:
			     树根
           /      |      \
         [1]     [2]     [3]
         / \     / \     / \
       [2] [3] [1] [3] [1] [2]
        |   |   |   |   |   |
       [3] [2] [3] [1] [2] [1]
遍历该树的所有子路径，即相当于输出该nums的所有排列组合，步骤为:
1.选择子节点，即从[1,2,3]中选择一个加入到排列中
2.如果排列中的数字没有达到nums的长度,则进入到决策树的下一层，即重复1操作。当排列中的数字达到nums的数量，则该排列完成，加入到结果中。进入3操作
3.将之前加入的数字重新放回到nums中，选择另一个子节点。继续1操作。当选择完后则回到上一层
*/

var res [][]int

func container(nums []int, num int) bool {
	for _, data := range nums {
		if data == num {
			return true
		}
	}
	return false
}

func backtrack(nums []int, path []int) {
	if len(path) == len(nums) { //回溯出口
		var c = make([]int, len(path))
		copy(c, path)
		res = append(res, c)
		return
	}

	for index := 0; index < len(nums); index++ {
		//过滤重复元素
		if container(path, nums[index]) {
			continue
		}
		//选择一个元素加入到path中
		path = append(path, nums[index])
		//进入到下一个根路径
		backtrack(nums, path)
		//取消之前的选择
		path = append(path[:len(path)-1])
	}
}

func permute(nums []int) [][]int {
	res = make([][]int, 0) //记录结果
	var path = make([]int, 0)
	backtrack(nums, path)
	return res
}

func main() {
}
