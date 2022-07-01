package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := make([]int, 0)
	array2 := array1

	fmt.Println(reflect.ValueOf(array1).Pointer() == reflect.ValueOf(array2).Pointer())
}
