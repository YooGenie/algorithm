package ect

import "fmt"

var tail = 0
var head = 0
var qu [max]int

func QueueMain() {
	queueInit()
	out()
	in(5)
	out()
	in(3)
	in(9)
	in(1)
	in(9)
	in(8)
	in(4)
	out()
}

func queueInit() {
	fmt.Println("stack 초기화 시켰습니다")
	head = 0
	tail = 0
}

func EmptyQueue() bool {
	return head == tail
}

func FullQueue() bool {
	return head == (tail+1)%max
}

func in(x int) {
	if FullQueue() {
		fmt.Println("queue 꽉 찼습니다.")
		return
	}
	qu[tail] = x
	fmt.Println(qu[tail], "를 삽입 했습니다")
	tail++
	if tail == max {
		tail = 0
	}
}

func out() {
	if EmptyQueue() {
		fmt.Println("queue 비어있습니다")
		return
	}
	result := qu[head]
	head++
	fmt.Println(result, "를 뺐습니다.")
	if head == max {
		head = 0
	}
}
