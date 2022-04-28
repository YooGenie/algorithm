package programmers

import "fmt"

func Al77884() {
	left := 24
	right := 27
	result := 0

	for i := left; i <= right; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count%2 == 0 {
			result += i
		} else {
			result -= i
		}
	}

	fmt.Println(result)
}
