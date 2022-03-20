package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al15552() {
	var l, a, b int
	reader := bufio.NewReader(os.Stdin)  //입력
	writer := bufio.NewWriter(os.Stdout) //출력

	defer writer.Flush()

	fmt.Fscanln(reader, &l)

	for i := 0; i < l; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}

}
