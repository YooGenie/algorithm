package programmers

import (
	"fmt"
)

func Al12903() {
	s := "qwert"
	result := ""
	half := len(s) / 2
	if len(s)%2 == 0 {
		result = string(s[half-1]) + string(s[half])
	} else {
		result = string(s[half])
	}

	fmt.Println(result)
}
