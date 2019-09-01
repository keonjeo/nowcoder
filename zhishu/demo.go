package main

import "fmt"

func main() {
	var i, num, tmp uint64

	_, err := fmt.Scan(&num)

	if err != nil {
		return
	}

	tmp = num

	for i = 2; i <= tmp; i++ {
		if num == 1 {
			break
		}

		if i == tmp {
			fmt.Printf("%d%s", i, " ")
			break
		}

		for num%i == 0 {
			num /= i
			fmt.Printf("%d%s", i, " ")
		}
	}
	fmt.Println()
}
