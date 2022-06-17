package ect

import "fmt"

func GreatestCommonFactor(m int, n int) (result int) {

	if n == 0 {
		fmt.Println("최소공약수 : ", m)
		return m
	}

	return GreatestCommonFactor(n, m%n)

}
