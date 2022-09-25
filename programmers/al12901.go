package programmers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Al12901() {
	a := 5
	b := 24

	month := fmt.Sprintf("%02s", strconv.Itoa(a))
	day := fmt.Sprintf("%02s", strconv.Itoa(b))

	s := "2016" + month + day
	t, _ := time.Parse("20060102", s)

	result := strings.ToUpper(t.Weekday().String())[0:3]

	fmt.Println(result)
}
