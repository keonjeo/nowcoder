package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		s := scanner.Text()

		for len(s) > 8 {
			fmt.Println(s[:8])
			s = s[8:]
		}

		arr := make([]rune, 0)
		for i := 0; i < 8-len(s); i++ {
			arr = append(arr, '0')
		}
		s += string(arr)

		fmt.Println(s)
	}
}
