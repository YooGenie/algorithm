package programmers

import "fmt"

func Al77484() {
	var lottos = []int{44, 1, 0, 0, 31, 25}
	var win_nums = []int{31, 10, 45, 1, 6, 19}
	winning := map[int]int{
		6: 1,
		5: 2,
		4: 3,
		3: 4,
		2: 5,
		1: 6,
		0: 6,
	}
	var result []int
	zeroCount := 0
	sameCount := 0

	for _, v := range lottos {
		if v == 0 {
			zeroCount++
		}
		for j := 0; j < len(win_nums); j++ {
			if v == win_nums[j] {
				sameCount++
			}

		}
	}

	min := sameCount
	max := sameCount + zeroCount

	result = append(result, winning[max])
	result = append(result, winning[min])

	fmt.Println(result)

}
