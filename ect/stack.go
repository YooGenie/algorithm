package ect

import "fmt"

const max = 5

var top = 0

var st [max]int

func StackMain() {
	stackInit()
	pop()
	push(5)
	pop()
	push(3)
	push(9)
	push(1)
	push(9)
	push(8)
	push(4)
	pop()
}

func stackInit() {
	fmt.Println("stack 초기화 시켰습니다")
	top = 0
}

func Empty() bool {
	return top == 0
}

func Full() bool {
	return top == max
}

func push(x int) {
	if Full() {
		fmt.Println("stack 꽉 찼습니다.")
		return
	}
	st[top] = x
	fmt.Println(st[top], "를 push 했습니다")
	top++
}

func pop() {
	if Empty() {
		fmt.Println("stack 비어있습니다")
		return
	}
	top--
	fmt.Println(st[top], "를 pop했습니다")
}
