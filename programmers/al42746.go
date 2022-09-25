package programmers

import (
	"fmt"
	"sort"
	"strconv"
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
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	temp := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]/10*len(str[i]) == nums[j]/10*len(str[j]) {
				if nums[i]%10*len(str[i]) < nums[j]%10*len(str[j]) {
					temp = nums[i]
					nums[i] = nums[j]
					nums[j] = temp
				}
			} else if nums[i]%10*len(str[i]) == nums[j]/10*len(str[j]) && nums[i]%10*len(str[i]) < nums[j]%10*len(str[j]) {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	result := ""
	for _, v := range nums {
		result += strconv.Itoa(v)
	}
	//result := strings.Join(nums, "")
	fmt.Println(result)
}
