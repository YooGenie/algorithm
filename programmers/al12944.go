package programmers

import (
	"fmt"
)

func Al12944() {
	var arr = []int{1, 2, 3, 4}
	var sum int
	var result float64

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	result = float64(sum) / float64(len(arr))
	fmt.Println(result)
}
