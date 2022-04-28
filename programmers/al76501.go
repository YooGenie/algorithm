package programmers

import (
	"fmt"
)

func Al76501() {
	var absolutes = []int{4, 7, 12}
	var signs = []bool{true, false, true}
	sum := 0

	for i := 0; i < len(absolutes); i++ {
		if signs[i] {
			sum += absolutes[i]
		} else {
			sum -= absolutes[i]
		}
	}

	fmt.Println(sum)
}
