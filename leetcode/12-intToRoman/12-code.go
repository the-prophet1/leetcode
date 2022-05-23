package main

var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hundreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var thousands = []string{"", "M", "MM", "MMM"}

func intToRoman(num int) string {
	return thousands[num/1000] + hundreds[num/100%10] + tens[num/10%10] + ones[num%10]
}
