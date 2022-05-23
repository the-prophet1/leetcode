package main

import "math"

/*
599. 两个列表的最小索引总和
假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。

你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。



示例 1:

输入: list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
输出: ["Shogun"]
解释: 他们唯一共同喜爱的餐厅是“Shogun”。
示例 2:

输入:list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["KFC", "Shogun", "Burger King"]
输出: ["Shogun"]
解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
*/

func findRestaurant(list1 []string, list2 []string) []string {
	m1 := make(map[string]int)
	res := make([]string, 0)
	currentIndex := math.MaxInt
	for i, s := range list1 {
		m1[s] = i
	}

	for i, s := range list2 {
		if index, ok := m1[s]; ok {
			if i+index < currentIndex {
				currentIndex = i + index
				res = res[0:0]
				res = append(res, s)
			} else if i+index == currentIndex {
				res = append(res, s)
			}
		}
	}
	return res
}

func main() {
	findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"})
}
