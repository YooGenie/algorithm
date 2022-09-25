package programmers

import (
	"fmt"
	"strings"
)

func Al12951() {
	s := "3people unFollowed me"

	result := strings.Title(strings.ToLower(s))

	fmt.Println(result)
}
