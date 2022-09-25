package programmers

import (
	"fmt"
)

func Al12950() {
	var arr1 = [][]int{{1, 2}, {2, 3}}
	var arr2 = [][]int{{3, 4}, {5, 6}}

	result := make([][]int, len(arr1), len(arr2))
	for i := 0; i < len(arr1); i++ {

		for j := 0; j < len(arr2[i]); j++ {
			result[i] = append(result[i], arr1[i][j]+arr2[i][j])
		}
	}

	fmt.Println(result)
}
