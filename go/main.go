package main

import (
	"fmt"
	"strings"
)

var pool string = "(){}[]"

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(Valid(input))
}

func Valid(s string) bool {
	tmp := ""

	for _, i := range s {
		if strings.Contains(pool, string(i)) {
			switch i {
			case '(':
				tmp = ")" + tmp
				break
			case '[':
				tmp = "]" + tmp
				break
			case '{':
				tmp = "}" + tmp
				break
			default:
				if tmp[0] == byte(i) {
					tmp = tmp[1:]
				} else {
					return false
				}
			}
		}
	}
	return true
}
