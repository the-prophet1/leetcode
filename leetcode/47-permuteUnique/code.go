package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

			     树根
           /      |      \
         [1]     [1]     [2]
         / \             /
       [1] [2]         [1]
        |   |           |
       [2] [1]         [1]


示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10

*/

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	var dfs func(nums []int, path []int)

	vis := make([]bool, len(nums))

	dfs = func(nums []int, path []int) {
		if len(path) == len(nums) {
			r := make([]int, len(path))
			copy(r, path)
			res = append(res, r)
			return
		}

		for i := 0; i < len(nums); i++ {
			//确保重复的元素是从左往右填写
			if vis[i] || (i > 0 && !vis[i-1] && nums[i] == nums[i-1]) {
				continue
			}

			path = append(path, nums[i])
			vis[i] = true
			dfs(nums, path)
			vis[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(nums, []int{})

	return res

}

func main() {
	nums := []int{1, 1, 2}

	fmt.Println(permuteUnique(nums))
}
