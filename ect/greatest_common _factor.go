package ect

import "fmt"

func GreatestCommonFactor(m int, n int) (result int) {

	if n == 0 {
		fmt.Println(m)
		return m
	}

	return GreatestCommonFactor(n, m%n)

}
