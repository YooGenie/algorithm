package programmers

import "fmt"

func Al86051() {
	var numbers = []int{1, 2, 3, 4, 6, 7, 8, 0}
	total := 0
	result := 0

	for i := 0; i < 10; i++ {
		total += i
	}

	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}

	result = total - result

	fmt.Println(result)
}
