package main

import (
	"fmt"
)

func Al10818() {
	var n int
	var str,min,max int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &str)
			if i==0 {
				min =str
				max =str
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
