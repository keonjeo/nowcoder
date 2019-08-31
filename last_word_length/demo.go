package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lastWordLength(input []string) int {
	return len(input[len(input)-1])
}

func main() {
	var scan = bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		input := scan.Text()
		words := strings.Split(input, " ")
		result := lastWordLength(words)
		fmt.Println(result)
	}
}
