package main

import "fmt"

func main() {

	for {
		var inputNumber int
		k, _ := fmt.Scanf("%d", &inputNumber)

		if k == 0 {
			break
		}

		var x int
		var content []int

		for i := 0; i < inputNumber; i++ {
			fmt.Scanf("%d", &x)
			content = append(content, x)
		}

		newContent := removeDuplicate(content)

		for i := 0; i < len(newContent); i++ {
			for j := i; j < len(newContent); j++ {
				if newContent[i] >= newContent[j] {
					newContent[i], newContent[j] = newContent[j], newContent[i]
				}
			}
		}

		for i := 0; i < len(newContent); i++ {
			fmt.Println(newContent[i])
		}
	}

}

func removeDuplicate(a []int) []int {
	var newA []int
	for _, n := range a {
		if contains(newA, n) {
			continue
		}
		newA = append(newA, n)
	}

	return newA
}

func contains(arr []int, c int) bool {
	for _, a := range arr {
		if a == c {
			return true
		}
	}
	return false
}
