package programmers

import (
	"fmt"
	"strconv"
)

func Al12947() {
	x := 10
	var arr []string
	sum := 0

	str := strconv.Itoa(x)

	for i := 0; i < len(str); i++ {
		arr = append(arr, string(str[i]))
		num, _ := strconv.Atoi(arr[i])
		sum += num

	}

	if x%sum == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}



