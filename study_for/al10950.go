package study_for

import (
	"fmt"
)

func Al10950() {
	var a, b, c int
	fmt.Scanln(&a)
	var arr []int
	for i := 0; i < a; i++ {
		fmt.Scanln(&b, &c)
		arr = append(arr, b+c)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
