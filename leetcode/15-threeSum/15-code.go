package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//	sort.Ints(nums)
	var negativeArray = make([]int, 0)    //负数切片
	var nonNegativeArray = make([]int, 0) //非负数切片
	zeroNum := 0

	var resultMap = make(map[string][]int)
	var ret = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			nonNegativeArray = append(nonNegativeArray, nums[i])
		} else {
			negativeArray = append(negativeArray, nums[i])
		}
		if nums[i] == 0 {
			zeroNum++
		}
	}
	if zeroNum >= 3 {
		s := fmt.Sprintln([]int{0, 0, 0})
		resultMap[s] = []int{0, 0, 0}
	}

	var negativeMap = make(map[int]int)
	var nonNegativeMap = make(map[int]int)

	for i := 0; i < len(negativeArray); i++ {
		negativeMap[negativeArray[i]] = i
	}

	for i := 0; i < len(nonNegativeArray); i++ {
		nonNegativeMap[nonNegativeArray[i]] = i
	}

	//2种取法，取2个非负数和1个负数，取2个负数和1个非负数

	//取1个负数和2个非负数 ,-2,1,1

	for i := 0; i < len(negativeArray); i++ {
		for j := 0; j < len(nonNegativeArray); j++ {
			data := 0 - nonNegativeArray[j] - negativeArray[i]
			index, err := nonNegativeMap[data]
			if err == true && index != j { //find success
				result := []int{negativeArray[i], nonNegativeArray[j], data}
				sort.Ints(result)
				s := fmt.Sprintln(result)
				resultMap[s] = result
			}
		}
	}

	//取1个非负数和2个负数 ,-2,1,1
	for i := 0; i < len(nonNegativeArray); i++ {
		for j := 0; j < len(negativeArray); j++ {
			data := 0 - negativeArray[j] - nonNegativeArray[i]
			index, err := negativeMap[data]
			if err == true && index != j { //find success
				result := []int{negativeArray[j], nonNegativeArray[i], data}
				sort.Ints(result)
				s := fmt.Sprintln(result)
				resultMap[s] = result
			}
		}
	}
	fmt.Println(resultMap)

	for _, value := range resultMap {
		ret = append(ret, value)
	}
	return ret
}

func threeSum1(nums []int) [][]int {
	var ans = make([][]int, 0)

	l := len(nums)
	if l < 3 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := l - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}
	}
	return ans
}

func main() {
	threeSum([]int{0, 0, 0, 0})
}
