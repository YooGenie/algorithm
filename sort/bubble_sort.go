package sort

import "fmt"

func BubbleSort() {
	var arr = []int{3, 5, 6, 1, 9}
	fmt.Println("정렬 전 : ", arr)
	min := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				min = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = min
			}
		}
	}

	fmt.Println("정렬 후 : ", arr)
}
