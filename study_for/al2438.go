package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al2438() {
	var l int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &l)
	for i := 1; i <= l; i++ {
		for j:=0 ; j<i;j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
