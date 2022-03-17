package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	res := 0
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	m['a'] = 4
	m['b'] = 9
	m['c'] = 40
	m['d'] = 90
	m['e'] = 400
	m['f'] = 900

	s = strings.Replace(s, "IV", "a", -1)
	s = strings.Replace(s, "IX", "b", -1)
	s = strings.Replace(s, "XL", "c", -1)
	s = strings.Replace(s, "XC", "d", -1)
	s = strings.Replace(s, "CD", "e", -1)
	s = strings.Replace(s, "CM", "f", -1)

	for i := 0; i < len(s); i++ {
		res += m[s[i]]
	}
	return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
