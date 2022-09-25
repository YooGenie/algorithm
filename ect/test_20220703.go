package ect

import "fmt"

func Test20220703() {
	//123 => 0
	//213 => 1
	var grade = []int{2, 1, 1, 3, 5, 2}
	result := 0

	for i := len(grade); i > 1; i-- {
		if grade[i-1] < grade[i-2] {
			result += grade[i-2] - grade[i-1]
			grade[i-2] = grade[i-1]
		}
	}
	fmt.Println(result)
}
