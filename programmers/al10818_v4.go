package programmers

import (
	"bufio"
	"fmt"
	"os"
)

func Al10818V4() {
	stdin := bufio.NewReader(os.Stdin)
	var n, min, max int
	_, err := fmt.Scan(&n)
	arr := make([]int, n)
	if err != nil {
		fmt.Println("error")
		stdin.ReadString('\n')
	}
	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d", &arr[i])
		if err != nil {
			fmt.Println("error")
			stdin.ReadString('\n')
		}
	}
	min = arr[0]
	max = arr[0]
	for i := 0; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println(min, max)
}
