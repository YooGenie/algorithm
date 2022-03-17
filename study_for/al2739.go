package study_for

import (
	"fmt"
)

func Al2739() {
	var a int
	fmt.Scanln(&a)
	for i := 1; i < 10; i++ {
		fmt.Println(a, "*", i, "=", a*i)
	}
}
