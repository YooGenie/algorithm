package ect

import "fmt"

func Array2nd(x, y int) {
	arr := make([][]int, x)
	arr2 := make([]int, y)

	count := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			arr2[j] = count
			count++
		}
		arr3 := make([]int, y)
		copy(arr3, arr2)
		arr[i] = arr3

	}
	fmt.Println("최종값 : ", arr)
}
