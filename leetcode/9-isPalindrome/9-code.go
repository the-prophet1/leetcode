package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversionNumber = 0
	var number = x

	for number != 0 { //数字反转
		reversionNumber *= 10
		reversionNumber += number % 10
		number /= 10
	}
	if reversionNumber == x {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(123212))
}
