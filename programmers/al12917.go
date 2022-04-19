package programmers

import (
	"fmt"
	"sort"
)

func Al12918() {
	var s = "Zbcdefg"

	var result string
	var arr []int

	for i := 0; i < len(s); i++ {
		arr = append(arr, int(s[i]))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i := 0; i < len(arr); i++ {
		result += string(arr[i])
	}
	fmt.Println(result)
}
