package programmers

import (
	"fmt"
)

func Al10818V3() {
	var n, min, max int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	min = arr[0]
	max = arr[0]
	for i := 0; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println(min, max)
}
