package programmers

import (
	"fmt"
	"sort"
)

func Al42748() {
	var array = []int{1, 5, 2, 6, 3, 7, 4}
	var commands = [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	var arr []int
	var result []int

	for i := 0; i < len(commands); i++ {

		if commands[i][0] == commands[i][1] {
			result = append(result, array[commands[i][0]])
		} else {
			arr = array[commands[i][0]-1 : commands[i][1]]
			sort.Ints(arr)
			result = append(result, arr[commands[i][2]-1])
		}
	}

	fmt.Println(result)
}
