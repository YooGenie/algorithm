package programmers

import (
	"fmt"
)

func Al10818V1() {
	var n, str, min, max int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &str)
		if i == 0 {
			min = str
			max = str
		}
		if min > str {
			min = str
		}
		if max < str {
			max = str
		}
	}
	fmt.Println(min, max)

}
