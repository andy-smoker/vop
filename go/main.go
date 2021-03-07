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

// Valid - проверяем валидность пары скобок
func Valid(s string) bool {
	// создаём для временного хранения зкрывающей скобки
	tmp := ""
	// перебераем символы в строке
	for _, i := range s {
		// определяем является ли символ одной из скобок
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
				// проверка закрывающей части скобки
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
