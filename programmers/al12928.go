package programmers

import (
	"fmt"
)

func Al12928() {
	n := 12
	result := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result += i
		}
	}

	fmt.Println(result)
}
