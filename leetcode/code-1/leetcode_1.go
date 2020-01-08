package main

//暴力双重循环
//时间复杂度：O(n^2),空间复杂度为O(1)
func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//使用hash表，并将target减去nums中的数，然后用哈希表查询结果
//时间复杂度:O(n),空间复杂度O(n)
func twoSumHash(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, num := range nums {
		hash[num] = index //此处的key需要是nums的值，而value是下标
	}

	for i := 0; i < len(nums); i++ {
		elem := target - nums[i]
		index, ok := hash[elem]
		if ok == true && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}
