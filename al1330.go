package main

import "fmt"

func Al1330()  {
	var a, b int
	fmt.Scanln(&a, &b)
	if a>b {
		fmt.Println(">")
	}else if a<b {
		fmt.Println("<")
	}else if a==b{
		fmt.Println("==")
	}
}
