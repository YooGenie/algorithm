package main

import (
	"fmt"
)

func Al10818V2(){
	var n,str, min, max int
	fmt.Scan(&n)
	var arr []int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &str)
			arr = append(arr, str)
		}
		min = arr[0]
		max = arr[0]
		for i := 0; i < n; i++ {
			if min > arr[i] {
				min = arr[i]
			}
			if max < arr[i] {
				max = arr[i]
			}
		}

	fmt.Println(min, max)
}
