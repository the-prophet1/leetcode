package main

import "fmt"

func simplifyPath(path string) string {
	if len(path) == 1 {
		return "/"
	}
	st := make([]string, 0)

	dir := make([]byte, 0)

	for i := 1; i < len(path); i++ {
		if path[i] != '/' {
			dir = append(dir, path[i])
		} else if string(dir) == ".." {
			dir = make([]byte, 0)
			if len(st) != 0 {
				st = st[:len(st)-1]
			}
		} else if string(dir) == "." || len(dir) == 0 {
			dir = make([]byte, 0)
			continue
		} else {
			st = append(st, string(dir))
			dir = make([]byte, 0)
		}
	}

	if string(dir) == "." || len(dir) == 0 {
		// nothing to do
	} else if string(dir) == ".." {
		if len(st) != 0 {
			st = st[:len(st)-1]
		}
	} else {
		st = append(st, string(dir))
	}

	res := ""
	for _, s := range st {
		res = res + "/" + s
	}

	if len(res) == 0 {
		return "/"
	}
	return res
}

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
