package study_if

import (
	"fmt"
)

func Al2480() {
	var one, two, three, sum, max, sameNum int
	fmt.Scanln(&one, &two, &three)

	if one == two && one == three && two == three {
		sum = 10000 + one*1000
	} else if one != two && one != three && two != three {
		switch {
		case one > two && one > three:
			max = one
		case two > one && two > three:
			max = two
		case three > one && three > two:
			max = three
		}
		sum = max * 100
	} else {
		switch {
		case one == two || one == three:
			sameNum = one
		case two == three:
			sameNum = two
		}
		sum = 1000 + sameNum*100
	}

	fmt.Println(sum)
}
