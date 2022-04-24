package programmers

import (
	"fmt"
	"strconv"
	"strings"
)

func Al81301() {
	s := "one4seveneight"
	var en = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var num = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < len(en); i++ {
		s = strings.ReplaceAll(s, en[i], num[i])
	}

	result, _ := strconv.Atoi(s)

	fmt.Println(result)

}
