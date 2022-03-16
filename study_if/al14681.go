package study_if

import "fmt"

func Al14681() {
	var x, y, result int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x > 0 && y > 0 {
		result = 1
	} else if x < 0 && y > 0 {
		result = 2
	} else if x < 0 && y < 0 {
		result = 3
	} else if x > 0 && y < 0 {
		result = 4
	}
	fmt.Println(result)
}
