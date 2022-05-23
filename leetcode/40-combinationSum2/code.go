package main

import (
	"fmt"
	"sort"
)

/*
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。



示例1:

输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例2:

输入: candidates =[2,5,2,1,2], target =5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <=candidates.length <= 100
1 <=candidates[i] <= 50
1 <= target <= 30

*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	var dfs func([]int, []int, int)
	dfs = func(nums []int, candidates []int, target int) {
		if target == 0 {
			s := make([]int, len(nums))
			copy(s, nums)
			res = append(res, s)
			return
		}
		if len(candidates) == 0 {
			return
		}

		if target-candidates[0] < 0 {
			return
		}

		//选择这个数字
		nums = append(nums, candidates[0])
		dfs(nums, candidates[1:], target-candidates[0])
		nums = nums[:len(nums)-1]

		//不选这个数字，即相同的数字也不选
		for len(candidates) > 1 {
			if candidates[0] == candidates[1] {
				candidates = candidates[1:]
			} else {
				break
			}
		}
		dfs(nums, candidates[1:], target)
	}

	dfs([]int{}, candidates, target)
	return res
}

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(nums, 8))
}
