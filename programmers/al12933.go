package programmers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Al12933() {
	var n int64 = 118372
	str := strconv.Itoa(int(n))
	arr := strings.Split(str, "")

	sort.Sort(sort.Reverse(sort.StringSlice(arr)))

	result, _ := strconv.ParseInt(strings.Join(arr,""),10,64)

	fmt.Println(result)
}
