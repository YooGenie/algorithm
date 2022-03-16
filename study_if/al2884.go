package study_if

import "fmt"

func Al2884() {
	var hour, minute int
	fmt.Scanln(&hour, &minute)

	if hour==0 && minute<45 && minute>=0{
		hour=23
		minute=60-(45-minute)
	}else {
		before45Minute := hour *60 + minute - 45
		minute = before45Minute%60
		hour = before45Minute/60
	}

	fmt.Println(hour,minute)
}

