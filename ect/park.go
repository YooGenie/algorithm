package ect

import (
	"fmt"
	"strconv"
)

func Solution(park []string, routes []string) []int {

	parkLen := len(park)
	//pakeAll := make([][]string, parkLen)
	start := findStart(park)

	for i := 0; i < len(routes); i++ {
		num, _ := strconv.Atoi(routes[i][2:])
		fmt.Println("숫자 : ", num, routes[i][2:])
		if routes[i][0:1] == "E" {
			fmt.Println("오른쪽")
			for j := 0; j < parkLen; j++ {
				if routes[start[1]][j:parkLen] != "X" && j < parkLen {
					start[1] = +num
				}
			}

		}
		if routes[i][0:1] == "S" {
			fmt.Println("아래쪽")
			for j := 0; j < parkLen; j++ {
				if routes[start[0]][j:parkLen] != "X" && j < parkLen {
					start[0] = +num
				}
			}
			//start[0] = +num
		}
		if routes[i][0:1] == "N" {
			fmt.Println("위쪽")
			for j := 0; j < parkLen; j++ {
				if routes[start[0]][j:parkLen] != "X" && j < parkLen {
					start[0] = +num
				}
			}
			//start[0] = -num
		}
		if routes[i][0:1] == "W" {
			fmt.Println("왼쪽")
			for j := 0; j < parkLen; j++ {
				if routes[start[1]][j:parkLen] != "X" && j < parkLen {
					start[1] = +num
				}
			}
			//start[1] = -num
		}
	}
	fmt.Println(start)
	return start
}

func findStart(park []string) []int {
	start := make([]int, 2)
	for i, v := range park {
		fmt.Println(v)
		for j, w := range v {
			if string(w) == "S" {
				start[0] = i
				start[1] = j
			}
		}
	}
	return start
}
