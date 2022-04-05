package programmers

import "fmt"

func Al12948() {
	var phoneNumber string
	phoneNumber = "01033334444"
	length := len(phoneNumber)
	result := ""
	for i := 0; i < length-4; i++ {
		result = result + "*"
	}

	result = result + phoneNumber[length-4:length]
	fmt.Println(result)
}
