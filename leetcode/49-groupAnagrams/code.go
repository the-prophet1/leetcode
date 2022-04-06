package main

import (
	"fmt"
	"sort"
)

/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]


提示：

1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

思路：将不同的异位词进行排序，使用hash进行判定是否相同
*/

type buf []byte

func (b buf) Len() int {
	return len(b)
}

func (b buf) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b buf) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		sb := buf(strs[i])
		sort.Sort(sb)
		s := string(sb)
		data, ok := m[s]
		if !ok {
			m[s] = append([]string(nil), strs[i])
		} else {
			m[s] = append(data, strs[i])
		}
	}
	for _, strings := range m {
		res = append(res, strings)
	}

	return res
}

func main() {
	strs := []string{""}
	fmt.Println(groupAnagrams(strs))
}
