package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	stack := list.New()
	length := len(s)

	for i := 0; i < length; i++ {
		if (s[i] == ']' || s[i] == ')' || s[i] == '}') && stack.Len() == 0 {
			return false
		}
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.PushBack(s[i])
		} else if s[i] == ')' && stack.Back().Value.(uint8) == '(' {
			stack.Remove(stack.Back())
		} else if s[i] == ']' && stack.Back().Value.(uint8) == '[' {
			stack.Remove(stack.Back())
		} else if s[i] == '}' && stack.Back().Value.(uint8) == '{' {
			stack.Remove(stack.Back())
		} else {
			return false
		}
	}
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("(])"))
}
