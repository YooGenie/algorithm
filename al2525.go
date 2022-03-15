package main

import "fmt"

func Al2525() {
	var hour, minute, time int
	fmt.Scanln(&hour, &minute)
	fmt.Scanln(&time)
	allMinute := 60 * 24
	afterMinute := hour*60 + minute + time

	if allMinute <= afterMinute {
		afterMinute = afterMinute - allMinute
		minute = afterMinute % 60
		hour = afterMinute / 60
	} else {
		minute = afterMinute % 60
		hour = afterMinute / 60
	}

	fmt.Println(hour, minute)
}
