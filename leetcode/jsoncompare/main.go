package main

import (
	"fmt"
	"reflect"
)

type Json map[string]interface{}

func CompareJson(json1 Json, json2 Json) bool {
	for k, v1 := range json1 {
		v2, ok := json2[k]
		if !ok {
			return false
		}
		if !compare(v1, v2) {
			return false
		}
	}
	return true
}

func compare(v1, v2 interface{}) bool {
	//value的类型不相同返回false
	if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
		return false
	}

	switch v1.(type) {
	case []interface{}:
		if !ArrayCompare(v1.([]interface{}), v2.([]interface{})) {
			return false
		}
	case Json:
		if !CompareJson(v1.(Json), v2.(Json)) {
			return false
		}
	default:
		if v1 != v2 {
			return false
		}
	}
	return true
}

func ArrayCompare(sli1 []interface{}, sli2 []interface{}) bool {
	if len(sli1) != len(sli2) {
		return false
	}

	for i, v1 := range sli1 {
		//value的类型不相同返回false
		if reflect.TypeOf(v1) != reflect.TypeOf(sli2[i]) {
			return false
		}
		if !compare(v1, sli2[i]) {
			return false
		}
	}
	return true
}

func main() {

	var json1 = Json{
		"a": "a",
		"b": "b",
		"m": Json{
			"a": append([]interface{}(nil), "data", 21, "2", "aa"),
			"b": 10,
		},
	}

	var json2 = Json{
		"a": "a",
		"b": "b",
		"m": Json{
			"a": append([]interface{}(nil), "data", 21, "2"),
			"b": 10,
		},
	}
	fmt.Println(CompareJson(json1, json2))
}
