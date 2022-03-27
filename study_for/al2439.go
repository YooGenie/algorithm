package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al2439() {
	var l int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &l)
	for i := 1; i <= l; i++ {
		for j := i; j < l; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := l; j < l+i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
