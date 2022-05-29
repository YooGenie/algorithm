package programmers

import (
	"fmt"
	"strings"
)

func Al12951() {
	s := "3people unFollowed me"

	result := ""

	for i := 0; i < len(s); i++ {
		if i == 0 {
			result = strings.ToUpper(string(s[i]))
		} else if string(s[i-1]) == " " {
			result += strings.ToUpper(string(s[i]))
		} else {
			result += strings.ToLower(string(s[i]))
		}
	}

	fmt.Println(result)
}
