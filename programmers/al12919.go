package programmers

import (
	"fmt"
	"strconv"
)

func Al12919() {
	var seoul = []string{"Jane", "Kim"}
	result := 0
	for i := 0; i < len(seoul); i++ {
		if seoul[i] == "Kim" {
			result = i
		}
	}

	fmt.Println("김서방은 " + strconv.Itoa(result) + "에 있다")
}
