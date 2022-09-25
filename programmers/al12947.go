package programmers

import (
	"fmt"
)

func Al12947() {
	x := 165
	sum := 0
	num := x

	for ; x >= 1; {
		sum += x % 10
		x = x / 10
	}

	if num%sum == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
