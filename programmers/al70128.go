package programmers

import "fmt"

func Al70128() {
	var a = []int{1, 2, 3, 4}
	var b = []int{-3, -1, 0, 2}
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}

	fmt.Println(result)
}
