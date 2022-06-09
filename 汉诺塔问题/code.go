package main

import "fmt"

type pedestal[T any] struct {
	ele *[]T
	pos int
	// pos代表了当前数组内元素数量
}

func hanoi[T any](pos int, A, B, C *pedestal[T]) {
	if pos > 0 {
		hanoi(pos-1, A, C, B)
		move(A, B)
		hanoi(pos-1, C, B, A)
	}
	// fmt.Println(A.ele,B.ele,C.ele)
}
func move[T any](A, B *pedestal[T]) {
	(*B.ele)[B.pos] = (*A.ele)[A.pos-1]
	B.pos++
	var t T
	(*A.ele)[A.pos-1] = t
	A.pos--
}
func main() {
	A := &pedestal[int]{&[]int{1, 2, 3, 4}, 4}
	B := &pedestal[int]{&[]int{0, 0, 0, 0}, 0}
	C := &pedestal[int]{&[]int{0, 0, 0, 0}, 0}
	hanoi(4, A, B, C)
	fmt.Println(A.ele,B.ele,C.ele)
	// &[0 0 0 0] &[1 2 3 4] &[0 0 0 0]
}
