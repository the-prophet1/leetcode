package main

import (
	"fmt"
	"sort"
)

//合并数组排序求中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, num := range nums2 {
		nums1 = append(nums1, num)
	}

	sort.Ints(nums1)

	length := len(nums1)
	if length%2 == 0 { //为偶数个
		index1 := length/2 - 1
		index2 := length / 2
		return float64(nums1[index1]+nums1[index2]) / 2
	} else {
		index := length / 2
		return float64(nums1[index])
	}
}

//使用归并排序的方式进行排序，时间复杂度为O(m+n)
func findMedianSortedArraysTopK(nums1 []int, nums2 []int) float64 {
	temp := make([]int, 0)
	var i, j, k int
	for i, j, k = 0, 0, 0; i < len(nums1) && j < len(nums2); k++ {
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		temp = append(temp, nums1[i])
		i++
	}
	for j < len(nums2) {
		temp = append(temp, nums2[j])
		j++
	}

	length := len(temp)
	if length%2 == 0 { //为偶数个
		index1 := length/2 - 1
		index2 := length / 2
		return float64(temp[index1]+temp[index2]) / 2
	} else {
		index := length / 2
		return float64(temp[index])
	}
}

func main() {
	nums1 := []int{2}
	nums2 := []int{1, 3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysTopK(nums1, nums2))
}
