package programmers

import "fmt"

func Al12931() {
	num := 100
	result := 0

	for ; num >= 10; {
		result = result + num%10
		num /= 10
	}

	fmt.Println(result + num)
}
