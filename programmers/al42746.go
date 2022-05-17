package programmers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Al42746() {
	var numbers = []int{3, 30, 34, 5, 9}

	var str []string
	var nums []int

	for _, v := range numbers {
		str = append(str, strconv.Itoa(v))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(str)))


	for _, v := range str {
		num ,_ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	for i:=0 ; i<len(nums); i++ {
		for j:=i+1 ;j<len(nums); j++  {
			if str[i][0] == str[j][0] && nums[i]%10 < nums[j]%10{
				sort.StringSlice(str).Swap(i, j)
				sort.IntSlice(nums).Swap(i, j)
			}
		}

	}



	result := strings.Join(str, "")
	fmt.Println(result)
}

