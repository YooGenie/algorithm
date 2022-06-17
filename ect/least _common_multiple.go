package ect

import "fmt"

func LeastCommonMultiple(m int, n int) (result int) {

	result = m * n / GreatestCommonFactor(m, n)

	fmt.Println("최소공배수 : ", result)

	return result
}
