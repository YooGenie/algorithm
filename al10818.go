package main

import (
	"fmt"
)

func Al10818() error {
	var n int
	var str, min, max int
	_, err := fmt.Scan(&n)
	if n <= 1 && n >= 1000000 {
		return err
	}
	var arr []int
	fmt.Println(n)
		for i := 0; i < n; i++ {
			_, err = fmt.Scanf("%d", &str)
			arr = append(arr, str)
			if str <= -1000000 && str >= 1000000 {
				return err
			}

		}
		fmt.Println(arr)
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
	return 	nil
}
