package main

import (
	"fmt"
	"strconv"
)

func Al2577(a int, b int, c int) {
	num := a * b * c
	str := strconv.Itoa(num)
	var slice []string
	r := [10]int{}

	for i := 0; i < len(str); i++ {
		slice = append(slice, str[i:i+1])
		b, _ := strconv.Atoi(slice[i])
		r[b] += 1
	}

	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
}
