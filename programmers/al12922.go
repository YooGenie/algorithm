package programmers

import (
	"fmt"
)

func Al12922() {
	n := 5
	result := ""

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			result += "박"
		} else {
			result += "수"
		}

	}

	fmt.Println(result)

}
