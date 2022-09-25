package programmers

import "fmt"

func Al12932() {
	var n int64
	n = 951236
	var result []int

	for ; n > 0; {
		result = append(result, int(n%10))
		n /= 10
	}

	fmt.Println(result)
}
