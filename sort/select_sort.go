package sort

import "fmt"

func SelectSort() {
	var arr = []int{3, 5, 6, 1, 9}
	fmt.Println("정렬 전 : ", arr)
	min := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				min = arr[j]
				arr[j] = arr[i]
				arr[i] = min
			}
		}
	}

	fmt.Println("정렬 후 : ", arr)
}
