package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al11021() {
	var l, a,b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()


	fmt.Fscanln(reader, &l)
	for i := 1; i <= l; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprint(writer, "Case #",i,": ",a+b,"\n")
	}
}