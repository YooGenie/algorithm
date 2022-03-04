package main

import "fmt"

func Al2753() {
	var y int
	fmt.Scanln(&y)
	if y%400 == 0 {
		fmt.Println(1)
	} else if y%4 == 0 && y%100 != 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
