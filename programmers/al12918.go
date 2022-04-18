package programmers

import (
	"fmt"
	"strconv"
)

func Al12918() {
	s := "a333"
	var result bool

	if len(s) == 4 || len(s) == 6 {
		if c, err := strconv.ParseInt(s, 10, 64); err != nil {
			result = false
		} else if c != 0 {
			result = true
		}
	}

	fmt.Println(result)
}
