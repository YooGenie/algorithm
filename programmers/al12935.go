package programmers

import "fmt"

func Al12935() {
	var arr = []int{4, 3, 2, 1}
	var result []int
	var min = arr[0]
	var index = 0

	if len(arr) == 0 || len(arr) == 1 {
		result = append(result, -1)
	}

	for i := 1; i < len(arr); i++ {
		if min >= arr[i] {
			min = arr[i]
			index = i
		}
	}

	for i := 0; i < index; i++ {
		result = append(result, arr[i])
	}

	for i := index + 1; i < len(arr); i++ {
		result = append(result, arr[i])
	}

	fmt.Println(result)
}
