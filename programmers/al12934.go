package programmers

import (
	"fmt"
)

func Al12934() {
	var result, i, n int64
	n = 121

	for i = 1; i <= n; i++ {
		if i*i == n {
			result = (i + 1) * (i + 1)
			break
		} else if n/i != n {
			result = -1
		}

	}
	fmt.Println(result)
}
