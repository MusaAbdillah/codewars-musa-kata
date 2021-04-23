package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello, playground")
	str := "abc def"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Fields(str))
	var a string
	testArray := strings.Fields(str)
	for _, v := range testArray {
		for i, rune := range v {
			if (i%2 == 0) && !unicode.IsSpace(rune) {
				a = a + strings.ToUpper(string(rune))
			} else if (i%2 != 0) && !unicode.IsSpace(rune) {
				a = a + strings.ToLower(string(rune))
			}
		}
		a = a + " "
	}

	fmt.Println(a)
	a = strings.TrimSuffix(a, " ")
	fmt.Println(len(a))
}
