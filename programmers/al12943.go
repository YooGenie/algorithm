package programmers

import "fmt"

func Al12943() {
	num := 16
	result := 0

	for ; num != 1 && result != -1; {
		if num%2 == 0 {
			num /= 2
		} else if num%2 == 1 {
			num = num*3 + 1
		}
		result += 1

		if result == 500 {
			result = -1
		}

	}
	fmt.Println(result)
}
