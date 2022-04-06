package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4

思路：
将区间进行排序，以左区间为基点从小到大排序。
比较nums[i - 1][1]与nums[i][0],如果nums[i - 1][1] 大于 nums[i][0]，说明nums[i]可以进行合并,否则不能进行合并
*/

type interval [][]int

func (i interval) Len() int {
	return len(i)
}

func (i interval) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}

func (i interval) Swap(a, b int) {
	i[a][0], i[b][0] = i[b][0], i[a][0]
	i[a][1], i[b][1] = i[b][1], i[a][1]
}

func merge(intervals [][]int) [][]int {
	ints := interval(intervals)
	res := make([][]int, 0)
	res = append(res, ints[0])

	sort.Sort(ints)

	for i := 1; i < len(ints); i++ {
		if res[len(res)-1][1] >= ints[i][0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], ints[i][1])
		} else {
			res = append(res, ints[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{{1, 4}, {0, 4}}
	fmt.Println(merge(nums))
}
