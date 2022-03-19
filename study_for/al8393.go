package study_for

import (
	"fmt"
)

func Al8393() {
	var a, sum int
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		sum += i
	}

	fmt.Println(sum)

}
