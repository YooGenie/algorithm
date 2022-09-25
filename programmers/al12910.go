package programmers

import (
	"fmt"
	"sort"
)

func Al12910() {
	var arr = []int{5, 9, 7, 10}

	divisor := 5
	var result []int

	for i := 0; i < len(arr); i++ {
		if arr[i]%divisor == 0 {
			result = append(result, arr[i])
		}
	}

	sort.Ints(result)
	if len(result) == 0 {
		result = append(result, -1)
	}

	fmt.Println(result)
}
