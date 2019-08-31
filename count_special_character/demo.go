package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstLine string
	var keyChar byte

	fmt.Scanf("%s", &firstLine)
	fmt.Scanf("%c", &keyChar)

	firstLine = strings.ToUpper(firstLine)

	var count int8
	for _, c := range firstLine {
		if isLowerCase(keyChar) {
			keyChar -= 32
		}

		if c == rune(keyChar) {
			count++
		}
	}

	fmt.Println(count)
}

func isLowerCase(a byte) bool {
	return a >= 97 && a <= 123
}
