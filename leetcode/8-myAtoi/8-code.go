package main

import (
	"fmt"
	"math"

	"strings"
)

func myAtoi(str string) int {
	var result int
	_, err := fmt.Sscanf(str, "%d", &result)
	if err != nil && strings.Contains(err.Error(), "value out of range") {
		if strings.Contains(str, "-") == true {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}

func main() {
	fmt.Println(myAtoi("   -2000000000000000000000000"))
}
