package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al10871() {
	var n, a, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &a)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		if a > x {
			fmt.Fprintf(writer, "%d ", x)
		}
	}

}
