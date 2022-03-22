package study_for

import (
	"bufio"
	"fmt"
	"os"
)

func Al2741() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader,&a)
	for i := 1; i <= a; i++ {
	fmt.Fprintln(writer,i)
	}
}
