package main

import (
	"fmt"
	"sort"
)

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

func main() {

	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 4))
}
