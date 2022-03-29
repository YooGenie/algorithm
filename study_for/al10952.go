package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al10952() {
	var a, b int

	reader := bufio.NewReader(os.Stdin)  //입력
	writer := bufio.NewWriter(os.Stdout) //출력
	defer writer.Flush()

	for i := 1; i > 0; i++ {
		fmt.Fscanln(reader, &a, &b)
		if a != 0 && b != 0 {
			fmt.Fprintln(writer, a+b)
			i++
		} else {
			i = -1
		}

	}

}
