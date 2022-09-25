package programmers

import (
	"fmt"
	"strings"
)

func Al12930() {
	s := "DREAm START"
	var arr = strings.Split(s, " ")
	result := ""

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if j%2 == 0 {
				result = result + strings.ToUpper(string(arr[i][j]))
			} else {
				result = result + strings.ToLower(string(arr[i][j]))
			}
		}

		if i < len(arr)-1 {
			result += " "
		}

	}

	fmt.Println(result)

}
