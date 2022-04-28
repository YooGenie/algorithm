package programmers

import "fmt"

func Al82612() {
	price := 3
	money := 20
	count := 4

	sum := 0
	result := 0

	for i := 1; i <= count; i++ {
		sum += price * i
	}
	if money < sum {
		result = sum - money
	}

	fmt.Println(result)
}