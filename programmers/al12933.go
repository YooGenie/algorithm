package programmers

import (
	"fmt"
	"sort"
	"strconv"
)

func Al12933() {
	var n int64 = 118372
	str := strconv.Itoa(int(n))
	var arr []string
	for i := 0; i < len(str); i++ {

		arr = append(arr, string(str[i]))

	}
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))

	resultStr := ""
	for _, v := range arr {
		resultStr += v
	}

	result, _ := strconv.ParseInt(resultStr, 10, 64)

	fmt.Println(result)
}
