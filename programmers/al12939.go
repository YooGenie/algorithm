package programmers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Al12939() {
	s := "1 2 3 4"
	num := 0

	var intArr []int
	strArr := strings.Split(s, " ")

	for _, v := range strArr {
		num, _ = strconv.Atoi(v)
		intArr = append(intArr, num)
	}

	sort.Ints(intArr)

	result := strconv.Itoa(intArr[0]) + " " + strconv.Itoa(intArr[len(strArr)-1])

	fmt.Println(result)
}
