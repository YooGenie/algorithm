package ect

import "fmt"

func FibonacciNumbers(n int) (result int) {


	if n==0 {
		return 0
	}else if n==1{
		return 1
	}

	fmt.Println(n)
	return FibonacciNumbers(n-1) + FibonacciNumbers(n-2)
}
