package programmers

import "fmt"

func Al87389() {
	n := 12
	x := 0

	for i := 1; n%i != 1; i++ {
		x = i + 1
	}

	fmt.Println(x)
}
