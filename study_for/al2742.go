package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al2742() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &a)
	for i := a; i >= 1; i-- {
		fmt.Fprintln(writer, i)
	}
}
