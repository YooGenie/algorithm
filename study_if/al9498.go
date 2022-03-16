package study_if

import "fmt"

func Al9498()  {
	var a int
	fmt.Scanln(&a)
	if a <=100 && a>=90 {
		fmt.Println("A")
	}else if a <90 && a>=80 {
		fmt.Println("B")
	}else if a <80 && a>=70 {
		fmt.Println("C")
	}else if a <70 && a>=60 {
		fmt.Println("D")
	}else {
		fmt.Println("F")
	}
}

