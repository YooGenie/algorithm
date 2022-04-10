package programmers

import "fmt"

func Al12937() {
	num := 16
	var result string

	if num%2 == 0 {
		result = "Even"
	} else if num%2 == 1 || num%2 == -1 {
		result = "Odd"
	}

	fmt.Println(result)
}
