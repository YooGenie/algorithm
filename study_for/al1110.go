package study_for

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Al1110() {
	var num, i, a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &num)
	first := num
	for i = 1; ; i++ {
		str := strconv.Itoa(num)
		if len(str) == 1 {
			a = 0
			b = num
		} else {
			a, _ = strconv.Atoi(string(str[0]))
			b, _ = strconv.Atoi(string(str[1]))
		}
		sum := a + b
		x := sum % 10
		num = (b * 10) + x
		if first == num {
			break
		}
	}
	fmt.Fprint(writer, i)
}
