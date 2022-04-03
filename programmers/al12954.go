package programmers

import "fmt"

func Al12954() {
	var x, n int
	var answer []int64

	fmt.Scan(&x, &n)

	for i := 1; i <= n; i++ {
		answer = append(answer, int64(x*i))
	}
	fmt.Println(answer)
}
